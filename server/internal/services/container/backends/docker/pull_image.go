/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package dockerbackend

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"regexp"
	"strconv"
	"sync"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/openorch/openorch/sdk/go/logger"
	"github.com/pkg/errors"
)

var progressRegex = regexp.MustCompile(`(\d+(\.\d+)?)([KMG]?B)/(\d+(\.\d+)?)([KMG]?B)`)

func (d *DockerBackend) pullImage(imageName string) error {
	d.imagePullGlobalMutex.Lock()

	imageMutex, exists := d.imagePullMutexes[imageName]
	if !exists {
		imageMutex = &sync.Mutex{}
		d.imagePullMutexes[imageName] = imageMutex
	}

	d.imagePullGlobalMutex.Unlock()

	imageMutex.Lock()
	defer imageMutex.Unlock()

	images, err := d.client.ImageList(context.Background(), image.ListOptions{
		All: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to list Docker images")
	}

	imageExists := false
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == imageName || tag == fmt.Sprintf("%v:latest", imageName) {
				imageExists = true
				break
			}
		}
		if imageExists {
			break
		}
	}

	if imageExists {
		return nil
	}

	logger.Info("Starting to pull image", slog.String("image", imageName))

	err = pullImageWithProgress(d.client, imageName)
	if err != nil {
		logger.Error("Failed to pull image",
			slog.String("image", imageName),
			slog.String("error", err.Error()),
		)
		return err
	}

	logger.Debug("Pulling image is done", slog.String("image", imageName))

	return nil
}

type PullStatus struct {
	Status   string `json:"status"`
	Progress string `json:"progress"`
	ID       string `json:"id"`

	// Fields set by OpenOrch

	ImageName string
}

func pullImageWithProgress(d *client.Client, imageName string) error {
	rc, err := d.ImagePull(context.Background(), imageName, image.PullOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to pull image")
	}
	defer func() {
		if err := rc.Close(); err != nil {
			logger.Error("Failed to close image pull response",
				slog.String("image", imageName),
				slog.String("error", err.Error()),
			)
		}
	}()

	decoder := json.NewDecoder(rc)
	for {
		var status PullStatus
		if err := decoder.Decode(&status); err == io.EOF {
			break
		} else if err != nil {
			logger.Error("Error pulling image",
				slog.String("error", err.Error()),
				slog.String("image", imageName),
			)
			return errors.Wrap(err, "Failed to decode image pull output")
		}
		status.ImageName = imageName
		logPullProgress(status)
	}

	return nil
}

// Example values:
// "status":"Downloading"
// "progress":"[================>                                  ]  722.2MB/2.219GB"
func logPullProgress(status PullStatus) {
	if status.Progress != "" {
		logger.Info("Pulling image progress",
			slog.String("status", status.Status),
			slog.String("progress", status.Progress),
			slog.String("imageName", status.ImageName),
		)
	} else {
		logger.Info("Pulling image",
			slog.String("status", status.Status),
			slog.String("imageName", status.ImageName),
		)
	}
}

func calculateProgressPercentage(progress string) float64 {
	matches := progressRegex.FindStringSubmatch(progress)
	if len(matches) < 7 {
		return 0
	}

	downloaded, _ := strconv.ParseFloat(matches[1], 64)
	downloadUnit := matches[3]
	total, _ := strconv.ParseFloat(matches[4], 64)
	totalUnit := matches[6]

	downloadedBytes := convertToBytes(downloaded, downloadUnit)
	totalBytes := convertToBytes(total, totalUnit)

	if totalBytes == 0 {
		return 0
	}
	return (downloadedBytes / totalBytes) * 100
}

// Docker seems to report sizes in base-10 (SI units) where 1 GB = 1000 MB, not 1024 MB.
const base = 1000

func convertToBytes(value float64, unit string) float64 {
	switch unit {
	case "KB":
		return value * base
	case "MB":
		return value * base * base
	case "GB":
		return value * base * base * base
	default:
		return value
	}
}
