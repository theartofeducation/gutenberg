package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var packagesFolder string

var root = &cobra.Command{
	Use:   "gutenberg",
	Short: "Version, Tag and Release your projects",
	Long:  `Handles all aspects of deploying a shared package for use in dependent projects`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	root.PersistentFlags().StringVarP(&packagesFolder, "packages-folder", "p", "components", "Relative path to packages folder")
}

// Execute run Gutenberg
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
