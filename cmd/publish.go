package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Publish = &cobra.Command{
	Use:   "publish",
	Short: "Publish your package(s) to the registry",
	Long: `This command will publish your package(s) to the configured package
				registry (typically either npm or GitHub).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Publishing your package")
	},
}
