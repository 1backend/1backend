package imageservice

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImageVariantCachePathStaysLocal(t *testing.T) {
	cacheRoot := t.TempDir()
	cs := &ImageService{imageCacheFolder: cacheRoot}
	hash := "0123456789abcdef0123456789abcdef01234567"

	cachePath := cs.getCachePath(hash)

	require.Equal(t, hash, filepath.Base(cachePath))
	rel, err := filepath.Rel(cacheRoot, cachePath)
	require.NoError(t, err)
	require.False(t, strings.HasPrefix(rel, ".."+string(filepath.Separator)))
	require.False(t, filepath.IsAbs(rel))
	require.DirExists(t, filepath.Dir(cachePath))
}

func TestImageServiceDoesNotDependOnDurableStorageProviders(t *testing.T) {
	forbiddenImports := []string{
		"cloud.google.com/go/storage",
		"github.com/1backend/1backend/server/internal/services/file",
	}
	forbiddenIdentifiers := map[string]struct{}{
		"CloudCacheProvider": {},
		"GCSProvider":        {},
		"StorageProvider":    {},
		"downloadStorage":    {},
		"uploadStorage":      {},
	}

	fset := token.NewFileSet()
	entries, err := os.ReadDir(".")
	require.NoError(t, err)

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(fset, name, nil, parser.ImportsOnly)
		require.NoError(t, err, name)
		for _, imp := range file.Imports {
			importPath := strings.Trim(imp.Path.Value, `"`)
			for _, forbidden := range forbiddenImports {
				require.NotEqual(t, forbidden, importPath, "%s must not import durable storage provider package %s", name, forbidden)
			}
		}

		file, err = parser.ParseFile(fset, name, nil, 0)
		require.NoError(t, err, name)
		ast.Inspect(file, func(node ast.Node) bool {
			ident, ok := node.(*ast.Ident)
			if !ok {
				return true
			}
			if _, forbidden := forbiddenIdentifiers[ident.Name]; forbidden {
				t.Fatalf("%s references durable storage identifier %q", name, ident.Name)
			}
			return true
		})
	}
}
