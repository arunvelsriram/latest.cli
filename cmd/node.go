package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arunvelsriram/latest.cli/pkg/common"
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
		client := common.NewRepository(&http.Client{})
		registry := node.NewNPMRegistry("https://registry.npmjs.org", client)
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
