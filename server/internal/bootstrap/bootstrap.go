package bootstrap

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	user "github.com/1backend/1backend/server/internal/services/user/types"
)

const (
	EntityPermit = "user-svc:permit"
	EntityEnroll = "user-svc:enroll"
	EntitySecret = "secret-svc:secret"
)

type Services struct {
	SavePermits func(context.Context, []user.PermitInput) error
	SaveEnrolls func(context.Context, []user.EnrollInput) error
}

type Summary struct {
	AppliedPermits int
	AppliedEnrolls int

	SkippedSecrets     int
	SkippedUnsupported int
	SkippedFiles       int
}

type meta struct {
	Entity string `json:"entity" yaml:"entity"`
}

type entityWrapper struct {
	Meta *meta `json:"_meta" yaml:"_meta"`
}

// Apply reads startup manifests from path and applies supported non-secret
// entities. Folder-level defaults come from _meta.yaml files and file-level
// defaults come from a top-level _meta key.
func Apply(ctx context.Context, path string, services Services) (*Summary, error) {
	if strings.TrimSpace(path) == "" {
		return &Summary{}, nil
	}

	stat, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "stat bootstrap path %s", path)
	}

	summary := &Summary{}
	if !stat.IsDir() {
		if err := applyFile(ctx, path, "", services, summary); err != nil {
			return nil, err
		}
		return summary, nil
	}

	metaByDir := map[string]string{}
	err = filepath.WalkDir(path, func(filePath string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			entity, err := entityForDir(path, filePath, metaByDir)
			if err != nil {
				return err
			}
			if entity != "" {
				metaByDir[filePath] = entity
			}
			return nil
		}

		if filepath.Base(filePath) == "_meta.yaml" || filepath.Base(filePath) == "_meta.yml" {
			return nil
		}
		if !isManifestFile(filePath) {
			summary.SkippedFiles++
			return nil
		}

		entity := metaByDir[filepath.Dir(filePath)]
		return applyFile(ctx, filePath, entity, services, summary)
	})
	if err != nil {
		return nil, err
	}

	return summary, nil
}

func applyFile(
	ctx context.Context,
	path string,
	inheritedEntity string,
	services Services,
	summary *Summary,
) error {
	doc, err := readYAML(path)
	if err != nil {
		return err
	}

	entity := fileEntity(doc)
	if entity == "" {
		entity = inheritedEntity
	}
	if entity == "" {
		entity = inferEntityFromPath(path)
	}
	if isSecretEntity(entity) || isSecretPath(path) {
		summary.SkippedSecrets++
		return nil
	}

	doc = stripFileMeta(doc)

	switch entity {
	case EntityPermit:
		items, err := decodeEntities[user.PermitInput](doc)
		if err != nil {
			return errors.Wrapf(err, "decode permits from %s", path)
		}
		if services.SavePermits == nil {
			return errors.Errorf("permit bootstrap unsupported for %s", path)
		}
		if err := services.SavePermits(ctx, items); err != nil {
			return errors.Wrapf(err, "apply permits from %s", path)
		}
		summary.AppliedPermits += len(items)
	case EntityEnroll:
		items, err := decodeEntities[user.EnrollInput](doc)
		if err != nil {
			return errors.Wrapf(err, "decode enrolls from %s", path)
		}
		if services.SaveEnrolls == nil {
			return errors.Errorf("enroll bootstrap unsupported for %s", path)
		}
		if err := services.SaveEnrolls(ctx, items); err != nil {
			return errors.Wrapf(err, "apply enrolls from %s", path)
		}
		summary.AppliedEnrolls += len(items)
	default:
		summary.SkippedUnsupported++
	}

	return nil
}

func entityForDir(root, dir string, metaByDir map[string]string) (string, error) {
	parentEntity := ""
	if dir != root {
		parentEntity = metaByDir[filepath.Dir(dir)]
	}

	for _, name := range []string{"_meta.yaml", "_meta.yml"} {
		metaPath := filepath.Join(dir, name)
		if _, err := os.Stat(metaPath); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return "", err
		}

		b, err := os.ReadFile(metaPath)
		if err != nil {
			return "", err
		}
		m := meta{}
		if err := yaml.Unmarshal(b, &m); err != nil {
			return "", errors.Wrapf(err, "decode %s", metaPath)
		}
		if m.Entity != "" {
			return m.Entity, nil
		}
	}

	if parentEntity != "" {
		return parentEntity, nil
	}
	return inferEntityFromDir(dir), nil
}

func readYAML(path string) (any, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "read %s", path)
	}
	var doc any
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return nil, errors.Wrapf(err, "decode yaml %s", path)
	}
	return normalizeYAML(doc), nil
}

func decodeEntities[T any](doc any) ([]T, error) {
	switch v := doc.(type) {
	case []any:
		items := make([]T, 0, len(v))
		for _, raw := range v {
			item, err := decodeEntity[T](stripFileMeta(raw))
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		return items, nil
	default:
		item, err := decodeEntity[T](v)
		if err != nil {
			return nil, err
		}
		return []T{item}, nil
	}
}

func decodeEntity[T any](doc any) (T, error) {
	var out T
	b, err := json.Marshal(doc)
	if err != nil {
		return out, err
	}
	if err := json.Unmarshal(b, &out); err != nil {
		return out, err
	}
	return out, nil
}

func fileEntity(doc any) string {
	m, ok := doc.(map[string]any)
	if !ok {
		return ""
	}
	raw, ok := m["_meta"]
	if !ok {
		return ""
	}
	b, err := json.Marshal(raw)
	if err != nil {
		return ""
	}
	wrapper := entityWrapper{}
	if err := json.Unmarshal([]byte(`{"_meta":`+string(b)+`}`), &wrapper); err != nil {
		return ""
	}
	if wrapper.Meta == nil {
		return ""
	}
	return wrapper.Meta.Entity
}

func stripFileMeta(doc any) any {
	m, ok := doc.(map[string]any)
	if !ok {
		return doc
	}
	copy := make(map[string]any, len(m))
	for k, v := range m {
		if k == "_meta" {
			continue
		}
		copy[k] = v
	}
	return copy
}

func normalizeYAML(v any) any {
	switch t := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(t))
		for k, v := range t {
			out[k] = normalizeYAML(v)
		}
		return out
	case map[any]any:
		out := make(map[string]any, len(t))
		for k, v := range t {
			out[jsonKey(k)] = normalizeYAML(v)
		}
		return out
	case []any:
		for i := range t {
			t[i] = normalizeYAML(t[i])
		}
		return t
	default:
		return v
	}
}

func jsonKey(v any) string {
	switch t := v.(type) {
	case string:
		return t
	default:
		b, _ := json.Marshal(t)
		return string(b)
	}
}

func isManifestFile(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".yaml", ".yml", ".json":
		return true
	default:
		return false
	}
}

func inferEntityFromPath(path string) string {
	switch filepath.Base(filepath.Dir(path)) {
	case "enrolls":
		return EntityEnroll
	case "permits":
		return EntityPermit
	case "secrets":
		return EntitySecret
	default:
		return ""
	}
}

func inferEntityFromDir(path string) string {
	switch filepath.Base(path) {
	case "enrolls":
		return EntityEnroll
	case "permits":
		return EntityPermit
	case "secrets":
		return EntitySecret
	default:
		return ""
	}
}

func isSecretEntity(entity string) bool {
	return entity == EntitySecret
}

func isSecretPath(path string) bool {
	for _, part := range strings.Split(filepath.Clean(path), string(filepath.Separator)) {
		if part == "secrets" {
			return true
		}
	}
	return false
}
