package cmd

import (
	"fmt"

	"github.com/arunvelsriram/latest.cli/pkg/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
