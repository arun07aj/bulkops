package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bkp",
	Short: "bkp is a CLI tool for performing bulk operations on files.",
	Long:  "BulkOps aka bkp is a CLI tool for performing bulk operations on files in the specified directory.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error while executing bkp '%s'\n", err)
		os.Exit(1)
	}
}
