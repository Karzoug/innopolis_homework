/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"08/internal/netdwn"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var (
	timeout       time.Duration
	workersNumber uint
)

func init() {
	rootCmd.Flags().UintVarP(&workersNumber, "number", "n", uint(runtime.NumCPU()), "workers number")
	rootCmd.Flags().DurationVarP(&timeout, "timeout", "t", 30*time.Second, "timeout")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "downloader [-n <workers number>] [-t <timeout>] <urls>...",
	Short: "A concurrent file downloader",
	Long:  `A concurrent file downloader`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := netdwn.Download(cmd.Context(),
			args,
			netdwn.NewConfig(int(workersNumber), timeout),
		); err != nil {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
