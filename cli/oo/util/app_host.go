package util

import (
	"context"
	"fmt"
	"sort"

	openapi "github.com/1backend/1backend/clients/go"
)

func AppIdsToHosts(ctx context.Context, apiClient *openapi.APIClient, appIds []string) (map[string]string, error) {
	deduped := map[string]struct{}{}
	for _, id := range appIds {
		deduped[id] = struct{}{}
	}
	appIds = make([]string, 0, len(deduped))
	for id := range deduped {
		appIds = append(appIds, id)
	}

	appIdToHost := map[string]string{}

	if len(appIds) > 0 {
		ids := make([]string, 0, len(appIds))
		ids = append(ids, appIds...)
		sort.Strings(ids)

		appsResp, _, listErr := apiClient.UserSvcAPI.
			ListApps(ctx).
			Body(openapi.UserSvcListAppsRequest{
				Ids: ids,
			}).
			Execute()
		if listErr != nil {
			return nil, fmt.Errorf("failed to list apps: %w", listErr)
		} else {
			for _, app := range appsResp.Apps {
				appIdToHost[app.Id] = app.Host
			}
		}
	}

	return appIdToHost, nil
}
