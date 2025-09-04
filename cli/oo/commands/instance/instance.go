/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package instance

import (
	"github.com/spf13/cobra"
)

func AddInstanceCommands(rootCmd *cobra.Command) {
	var instanceCmd = &cobra.Command{
		Use:     "instance",
		Aliases: []string{"i", "instances"},
		Short:   "Manage service instances",
	}

	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all environments",
		RunE:    List,
	}

	var deleteCmd = &cobra.Command{
		Use:     "remove [instanceId]",
		Aliases: []string{"rm"},
		Short:   "Remove service instance",
		Args:    cobra.ExactArgs(1),
		RunE:    Remove,
	}

	instanceCmd.AddCommand(listCmd)
	instanceCmd.AddCommand(deleteCmd)

	rootCmd.AddCommand(instanceCmd)
}
