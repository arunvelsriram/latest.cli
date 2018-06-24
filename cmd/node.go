package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arunvelsriram/latest.cli/pkg/node"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:     "node <module-name>",
	Short:   "Get latest version of a node module",
	Aliases: []string{"n"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		registry := node.NewNPMRegistry("https://registry.npmjs.org", &http.Client{})
		version, err := registry.LatestVersion(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
