package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(Tag)
}

// Tag executes commands for applying a tag to the package repository
// using the version from the package.json file
var Tag = &cobra.Command{
	Use:   "tag",
	Short: "Apply a version tag to your repository",
	Long: `This command will tag your repository with the version number extracted
				from your package.json file. This version tag will be used for publishing
				the package to the registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tagging your repository")
	},
}
