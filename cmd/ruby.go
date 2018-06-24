package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arunvelsriram/latest.cli/pkg/ruby"
	"github.com/spf13/cobra"
)

var rubyCmd = &cobra.Command{
	Use:     "ruby <gem-name>",
	Short:   "Get latest version of a ruby gem",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		repo := ruby.NewGemRepository("https://rubygems.org/api/v1/gems", &http.Client{})
		version, err := repo.LatestVersion(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(rubyCmd)
}
