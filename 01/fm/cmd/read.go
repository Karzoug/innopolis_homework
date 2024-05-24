package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const maxReadSize int64 = 1 << 18 // 256 KiB

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:     "read filename",
	Short:   fmt.Sprintf("Read the file and print its content (limited to %d KiB)", maxReadSize/1024),
	Example: "  fm read README.txt",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(strings.TrimSpace(args[0]))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer f.Close()

		n, err := io.CopyN(os.Stdout, f, maxReadSize)
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if n == maxReadSize {
			fmt.Fprint(os.Stderr, "\n\n(file too large to print)")
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
