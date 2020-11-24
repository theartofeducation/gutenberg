package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theartofeducation/gutenberg/core"
	"github.com/ttacon/chalk"
)

func init() {
	root.AddCommand(Version)
}

// Version executes commands for incrementing the package version
// number in package.json and updating the package CHANGELOG.md
var Version = &cobra.Command{
	Use:   "version",
	Short: "Increment the package version",
	Long: `This command will increment your package version (major, minor, patch)
and update your CHANGELOG.md file. These changes will NOT be committed
to your repository, and you will be able to review the changes made before
committing them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(chalk.Cyan, "📦 Incrementing version and updating CHANGELOG for project package(s)", chalk.Reset)
		core.RunTaskForPackages(packages, "version")
	},
}
