package sdk

import (
	"os"
	"path"

	"github.com/pkg/errors"
)

const onebackendFolder = ".1backend"

type HomeDirOptions struct {
	Test bool
}

func HomeDir(options HomeDirOptions) (string, error) {
	var (
		homeDir string
		err     error
	)
	if options.Test {
		homeDir, err = os.MkdirTemp("", "1backend-")
		if err != nil {
			return "", errors.Wrap(err,
				"homedir creation failed",
			)
		}
	} else {
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", errors.Wrap(err, "homedir creation failed")
		}
		homeDir = path.Join(homeDir, onebackendFolder)
	}

	return homeDir, nil
}
