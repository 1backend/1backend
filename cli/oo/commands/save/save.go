/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package save

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	header  = color.New(color.FgCyan, color.Bold).SprintFunc()
	success = color.New(color.FgGreen, color.Bold).SprintFunc()
	warn    = color.New(color.FgYellow, color.Bold).SprintFunc()
	fail    = color.New(color.FgRed, color.Bold).SprintFunc()
	idStyle = color.New(color.FgHiBlack, color.Bold, color.Underline).SprintFunc()
)

type Entity struct {
	Meta *Meta `yaml:"_meta"`
}

type Meta struct {
	// Entity identifies the type of metadata. It can be used to link to external
	// metadata sources, or make them unnecessary when the entity is handled
	// directly (e.g. in hardcoded cases).
	Entity    string     `yaml:"entity,omitempty" example:"secret-svc:secret"`
	Version   string     `yaml:"version,omitempty"`
	Transport *Transport `yaml:"transport,omitempty"` // http|grpc; default http
}

type Transport struct {
	// Endpoint is the target URL or path. Required.
	Endpoint string `yaml:"endpoint"`

	// Method is the HTTP verb used for the request.
	// Supported: PUT, POST, PATCH. Default: PUT.
	Method string `yaml:"method,omitempty"`

	// Body defines how the payload is wrapped or transformed before sending.
	// - If set to a string (e.g. "data"), the payload is wrapped: {"data": body}.
	// - If set to a JSON object containing "$", "$" is replaced with the payload.
	//   Simple Example:   "key" →
	//                     { "key": "example payload of any type" }
	//   Complex Example:  {"data": "$", "meta": {"id": "123"}} →
	//                     "data": "example payload of any type", "meta": {"id": "123"}}
	Body any `yaml:"body,omitempty"`

	// Array indicates whether the endpoint expects an array payload. Default: false.
	Array bool `yaml:"array,omitempty"`

	// ContentType specifies the MIME type of the request body. Default: application/json.
	ContentType string `yaml:"contentType,omitempty"`
}

func AddSaveCommand(rootCmd *cobra.Command) {
	var saveCmd = &cobra.Command{
		Use:     "save",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Save(cmd, args)
		},
		Short: "Save a file or all files in a directory recursively",
		Long:  `Save a file or all files in a directory recursively. Files may have a _meta tag or inherit defaults from _meta.yaml in the folder.`,
	}
	rootCmd.AddCommand(saveCmd)
}

// Save [file | directory]
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	if len(args) == 0 {
		return errors.New("path required")
	}
	path := args[0]

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Errorf("path not found: %s", path)
	} else if err != nil {
		return err
	}

	if stat.IsDir() {
		return processDir(ctx, url, token, path)
	}
	return handleFileWithMeta(ctx, url, token, path, nil)
}

// processDir loads _meta.yaml once per directory, walks files, and applies defaults
func processDir(ctx context.Context, baseURL, token, dir string) error {
	var folderMeta *Meta

	return filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Println(warn("Entering folder:"), p)

			metaPath := filepath.Join(p, "_meta.yaml")

			_, err := os.Stat(metaPath)
			if err == nil || !os.IsNotExist(err) {
				b, err := ioutil.ReadFile(metaPath)
				if err != nil {
					return errors.Wrap(err, "cannot read _meta.yaml")
				}
				fmt.Println(warn("Loading folder meta from"), metaPath)

				err = yaml.Unmarshal(b, &folderMeta)
				if err != nil {
					return errors.Wrap(err, "invalid _meta.yaml")
				}
			}

			return nil
		}

		if filepath.Base(p) == "_meta.yaml" {
			return nil
		}

		return handleFileWithMeta(ctx, baseURL, token, p, folderMeta)
	})
}

// handleFileWithMeta merges file-level meta with folder defaults
func handleFileWithMeta(ctx context.Context, baseURL, token, filePath string, folderMeta *Meta) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	var wrapper Entity
	_ = yaml.Unmarshal(data, &wrapper)

	meta := wrapper.Meta
	if meta == nil {
		meta = &Meta{}
	}
	if meta.Transport == nil {
		meta.Transport = &Transport{}
	}

	if folderMeta != nil {
		if folderMeta.Transport != nil && folderMeta.Transport.Endpoint != "" {
			if meta.Transport.Endpoint == "" {
				meta.Transport.Endpoint = folderMeta.Transport.Endpoint
			}
			if meta.Transport.Method == "" {
				meta.Transport.Method = folderMeta.Transport.Method
			}
			if meta.Transport.Body == "" {
				meta.Transport.Body = folderMeta.Transport.Body
			}
			if meta.Transport.ContentType == "" {
				meta.Transport.ContentType = folderMeta.Transport.ContentType
			}
			if !meta.Transport.Array && folderMeta.Transport.Array {
				meta.Transport.Array = true
			}
		}

		if meta.Entity == "" && folderMeta.Entity != "" {
			meta.Entity = folderMeta.Entity
		}
	}

	if meta.Transport.Endpoint == "" {
		switch meta.Entity {
		case "secret-svc:secret":
			meta = &secretMeta
		case "user-svc:permit":
			meta = &permitMeta
		case "proxy-svc:route":
			meta = &routeMeta
		default:
			return errors.Errorf("unknown hardcoded entity %s in %s", meta.Entity, filePath)
		}
	}

	return handleEntity(ctx, baseURL, token, *meta, data, filePath)
}

func handleEntity(
	ctx context.Context,
	baseURL, token string,
	meta Meta,
	data []byte,
	filePath string,
) error {
	return handleDynamic(ctx, filePath, baseURL, token, meta, data)
}

func handleDynamic(
	ctx context.Context,
	filePath,
	baseURL, token string,
	meta Meta,
	fileBytes []byte,
) error {
	var doc map[string]interface{}
	if err := yaml.Unmarshal(fileBytes, &doc); err != nil {
		return errors.Wrap(err, "invalid yaml")
	}
	delete(doc, "_meta")

	var body interface{} = doc

	transport := meta.Transport

	if transport.Array {
		switch b := body.(type) {
		case []interface{}:
		default:
			body = []interface{}{b}
		}
	}

	switch b := transport.Body.(type) {
	case string:
		body = map[string]interface{}{b: body}
	default:
		return errors.Errorf("only string body wrapping is supported for now, not %T", b)
	}

	fmt.Printf("%s %s=%s %s=%s %s=%s... ", success("Saving"),
		header("file"), idStyle(filePath),
		header("endpoint"), idStyle(transport.Endpoint),
		header("method"), idStyle(transport.Method),
	)

	ct := transport.ContentType
	if ct == "" {
		ct = "application/json"
	}

	var payload []byte
	var err error
	switch ct {
	case "application/json":
		payload, err = json.Marshal(body)
		if err != nil {
			return errors.Wrap(err, "marshal json body")
		}
	case "application/yaml", "text/yaml":
		payload, err = yaml.Marshal(body)
		if err != nil {
			return errors.Wrap(err, "marshal yaml body")
		}
	default:
		return errors.Errorf("unsupported contentType: %s", ct)
	}

	method := strings.ToUpper(strings.TrimSpace(transport.Method))
	if method == "" {
		method = "PUT"
	}

	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(baseURL, "/")+meta.Transport.Endpoint, bytes.NewReader(payload))
	if err != nil {
		return errors.Wrap(err, "build request")
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", ct)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request failed")
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println(fail(fmt.Sprintf("dynamic save failed: %s: %s", resp.Status, string(bodyBytes))))
	}

	fmt.Println(success("ok"))
	return nil
}
