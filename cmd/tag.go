package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
