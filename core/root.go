package core

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log *logrus.Logger
var workingDirectory string

var root = &cobra.Command{
	Use:   "gutenberg",
	Short: "Version, Tag and Release your projects",
	Long:  `Handles all aspects of deploying a shared package for use in dependent projects`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	// root.PersistentFlags().StringVarP(&packagesFolder, "packages-folder", "p", "", "Relative path to packages folder")
}

// Execute runs Gutenberg. Who runs Gutenberg? Masterblaster runs Gutenberg.
func Execute(logger *logrus.Logger, directory string) {
	log = logger
	workingDirectory = directory

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
