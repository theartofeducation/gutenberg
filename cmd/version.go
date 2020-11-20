package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "Increment the package version",
	Long: `This command will increment your package version (major, minor, patch)
				and update your CHANGELOG.md file. These changes will NOT be committed
				to your repository, and you will be able to review the changes made before
				committing them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Incrementing package version and updating CHANGELOG.md")
	},
}
