package redirect

import (
	"net/http"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [--id --target | filePath | dirPath]
func Save(
	cmd *cobra.Command,
	args []string,
	id string,
	target string,
	statusCode int,
) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	var redirects []RedirectInput
	if id != "" && target != "" {
		redirects = []RedirectInput{
			{
				Id:         id,
				Target:     target,
				StatusCode: statusCode,
			},
		}
	} else {
		if len(args) == 0 {
			return errors.New("file or directory path required unless --id and --target are provided")
		}

		redirects, err = util.CollectFromPath[RedirectInput](args[0], "redirects")
		if err != nil {
			return err
		}
	}

	return doJSON(ctx, url, token, http.MethodPut, "/proxy-svc/redirects", SaveRedirectsRequest{
		Redirects: redirects,
	}, nil)
}
