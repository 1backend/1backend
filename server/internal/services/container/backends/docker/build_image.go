package dockerbackend

import (
	"archive/tar"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"

	"github.com/openorch/openorch/sdk/go/logger"
	container "github.com/openorch/openorch/server/internal/services/container/types"
)

func (dm *DockerBackend) BuildImage(req container.BuildImageRequest) (*container.BuildImageResponse, error) {
	ctx := context.Background()

	tarBuffer, err := createTarFromContext(req.ContextPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create build context tar")
	}

	dockerfilePath := req.DockerfilePath
	if dockerfilePath == "" {
		dockerfilePath = "Dockerfile"
	}

	options := types.ImageBuildOptions{
		Tags:           []string{req.Name},
		Dockerfile:     dockerfilePath,
		Remove:         true, // Remove intermediate containers
		ForceRemove:    true,
		SuppressOutput: false,
	}

	imageBuildResponse, err := dm.client.ImageBuild(ctx, tarBuffer, options)
	if err != nil {
		return nil, errors.Wrap(err, "image build failed")
	}
	defer imageBuildResponse.Body.Close()

	// Stream the build output to logs
	if err := streamBuildOutput(imageBuildResponse.Body); err != nil {
		return nil, errors.Wrap(err, "build failed")
	}

	return &container.BuildImageResponse{}, nil
}

func createTarFromContext(sourceDir string) (io.Reader, error) {
	pr, pw := io.Pipe()
	tw := tar.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer tw.Close()

		err := filepath.Walk(
			sourceDir,
			func(file string, fi os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// Preserve the directory structure in the tar file
				if fi.IsDir() {
					// Add directory header (empty, no contents)
					header, err := tar.FileInfoHeader(fi, fi.Name())
					if err != nil {
						return err
					}
					header.Name = filepath.ToSlash(file[len(sourceDir):])
					if err := tw.WriteHeader(header); err != nil {
						return err
					}
					return nil // Skip adding files for directories, but still write header
				}

				header, err := tar.FileInfoHeader(fi, fi.Name())
				if err != nil {
					return err
				}

				header.Name = filepath.ToSlash(file[len(sourceDir):])
				if err := tw.WriteHeader(header); err != nil {
					return err
				}

				if fi.Mode().IsRegular() {
					f, err := os.Open(file)
					if err != nil {
						return err
					}
					defer f.Close()
					_, err = io.Copy(tw, f)
					return err
				}
				return nil
			},
		)

		if err != nil {
			pw.CloseWithError(err)
		}
	}()

	return pr, nil
}

// JSON structure for Docker build output
type BuildOutput struct {
	Stream      string `json:"stream"`
	ErrorDetail struct {
		Message string `json:"message"`
	} `json:"errorDetail"`
	Error string `json:"error"`
}

// streamBuildOutput reads the output and detects errors
func streamBuildOutput(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse JSON line
		var output BuildOutput
		if err := json.Unmarshal([]byte(line), &output); err != nil {
			logger.Error("Failed to parse line as JSON",
				slog.String("line", line),
				slog.Any("error", err),
			)
			continue
		}

		// Print stream content
		if output.Stream != "" {
			logger.Info(output.Stream)
		}

		// Check for errors
		if output.Error != "" || output.ErrorDetail.Message != "" {
			logger.Error("Build failed",
				slog.String("error", output.Error))

			return errors.New(output.ErrorDetail.Message)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading stream: %w", err)
	}

	return nil
}
