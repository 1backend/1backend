/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	dapper "github.com/openorch/openorch/dapper/app"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "dapper"}
	var folder string

	// Extract CLI parameters
	params, remainingArgs := extractParams(os.Args)
	os.Args = remainingArgs

	var anon bool
	var retry int
	var retrySleepDuration string

	var runCmd = &cobra.Command{
		Use:   "run [config file]",
		Short: "Run app",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appFilePath := args[0]
			run(appFilePath, params, anon, retry, retrySleepDuration)
		},
	}

	// Global flag across all subcommands
	rootCmd.PersistentFlags().StringVarP(&folder, "folder", "f", ".", "directory containing configuration files")

	runCmd.Flags().BoolVar(&anon, "anon", false, "Run in anonymous mode")
	runCmd.Flags().IntVar(&retry, "retry", 0, "How many times to retry in case of failure")
	runCmd.Flags().StringVar(&retrySleepDuration, "retry-delay", "3s", "Delay between retries")

	rootCmd.AddCommand(runCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func extractParams(args []string) (map[string]string, []string) {
	params := make(map[string]string)
	var newArgs []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--var-") {
			split := strings.SplitN(arg[2:], "=", 2)
			if len(split) == 2 {
				key := strings.TrimPrefix(split[0], "var-")
				value := split[1]
				params[key] = value
			}
		} else {
			newArgs = append(newArgs, arg)
		}
	}
	return params, newArgs
}

func run(appFilePath string, params map[string]string, anon bool, retry int, retrySleep string) {
	cm := dapper.NewConfigurationManagerFromSource()
	app, err := cm.LoadAppConfiguration(appFilePath)
	if err != nil {
		log.Fatalf("Failed to load app file: %v", err)
	}

	var retryDelay time.Duration

	retryDelay, err = time.ParseDuration(retrySleep)
	if err != nil {
		fmt.Printf("Retry delay cannot be parsed, going with 1s")
		retryDelay = time.Second
	}

	retryLog := fmt.Sprintf("Retries: %v", retry)
	if retry > 0 {
		retryLog += fmt.Sprintf(", Retry Delay: %v", retryDelay)
	}

	fmt.Printf("%v, Parameters:\n", retryLog)
	if len(params) == 0 {
		fmt.Println("   None")
	}
	for key, value := range params {
		fmt.Printf("   %v=%v\n", key, value)
	}

	i := 0

	for {
		cont, err := cm.Run(app, params, anon)
		if err != nil {
			fmt.Printf("Failed to resolve feature dependencies: %v\n", err)
			if cont != nil && cont.RebootRequired {
				fmt.Printf("A restart is required to fix this!\n")
			}
			if i >= retry {
				os.Exit(1)
			} else {
				fmt.Printf("Retrying in %v", retryDelay)
				i++
				time.Sleep(retryDelay)
			}
		} else {
			os.Exit(0)
		}
	}
}
