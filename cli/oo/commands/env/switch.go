/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package env

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/spf13/cobra"
)

// Select [envShortName]
func Select(cmd *cobra.Command, args []string) error {
	conf, err := util.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	shortName := args[0]

	_, ok := conf.Environments[shortName]
	if !ok {
		return fmt.Errorf("env %v does not exist", shortName)
	}

	conf.SelectedEnvironment = shortName
	return util.SaveConfig(conf)
}
