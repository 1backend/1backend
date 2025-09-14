package route

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [--id --target | filePath | dirPath]
func Save(
	cmd *cobra.Command,
	args []string,
	id string,
	target string,
) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	// Case 1: Flags-based route
	if id != "" && target != "" {
		_, _, err := cf.Client(client.WithToken(token)).
			ProxySvcAPI.SaveRoutes(ctx).
			Body(
				openapi.ProxySvcSaveRoutesRequest{
					Routes: []openapi.ProxySvcRouteInput{
						{
							Id:     id,
							Target: &target,
						},
					},
				},
			).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to save route")
		}
		return nil
	}

	// Case 2: File or directory-based routes
	path := args[0]

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
	} else if err != nil {
		return errors.Wrap(err, "error checking path")
	}

	var routes []openapi.ProxySvcRouteInput
	fileCount := 0

	if stat.IsDir() {
		// Handle directory: Iterate over files and collect routes
		err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
			}
			if !info.IsDir() {
				// Collect routes from each file in the directory
				fileCount++
				fileRoutes, err := extractRoutesFromFile(filePath)
				if err != nil {
					return err
				}
				routes = append(routes, fileRoutes...)
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		// Handle single file
		fileCount++
		fileRoutes, err := extractRoutesFromFile(path)
		if err != nil {
			return err
		}
		routes = append(routes, fileRoutes...)
	}

	// Make a single API call to save all routes
	_, _, err = cf.Client(client.WithToken(token)).
		ProxySvcAPI.SaveRoutes(ctx).
		Body(openapi.ProxySvcSaveRoutesRequest{
			Routes: routes,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save routes")
	}

	return nil
}

// extract one or more routes from a file
func extractRoutesFromFile(filePath string) ([]openapi.ProxySvcRouteInput, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to read file at '%v'", filePath))
	}

	// Determine whether the file contains single or multiple routes
	var routes []openapi.ProxySvcRouteInput

	// Unmarshal as list first (multiple routes)
	if err := yaml.Unmarshal(data, &routes); err != nil {
		// If unmarshalling to list fails, attempt unmarshalling as single route
		var singleRoute openapi.ProxySvcRouteInput
		if err := yaml.Unmarshal(data, &singleRoute); err != nil {
			return nil, errors.Wrap(
				err,
				fmt.Sprintf(
					"failed to parse routes file at '%v' as single or multiple routes",
					filePath,
				),
			)
		}
		routes = append(routes, singleRoute)
	}

	return routes, nil
}
