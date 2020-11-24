package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/theartofeducation/gutenberg/core"
)

var (
	packagesFolder string   // The folder that contains all the packages to execute the actions on.
	packages       []string // List of folders the actions will be executed on.
)

var root = &cobra.Command{
	Use:   "gutenberg",
	Short: "Version, Tag and Release your projects",
	Long:  `Handles all aspects of deploying a shared package for use in dependent projects`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var packageFolder string
		var err error

		currentWorkingDirectory, err := os.Getwd()

		if packagesFolder != "" {
			packageFolders := core.GetPackageFolders(packagesFolder)
			for _, pf := range packageFolders {
				packageFolder, err = filepath.Abs(path.Join(currentWorkingDirectory, packagesFolder, pf.Name()))

				if err != nil {
					fmt.Printf("ERROR listing package folders: %v", err)
					os.Exit(1)
				}

				packages = append(packages, packageFolder)
			}
		} else {
			packageFolder, err = filepath.Abs(currentWorkingDirectory)

			if err != nil {
				fmt.Printf("ERROR listing package folder: %v", err)
				os.Exit(1)
			}

			packages = append(packages, packageFolder)
		}

		if err != nil {
			fmt.Println(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	root.PersistentFlags().StringVarP(&packagesFolder, "packages-folder", "p", "", "Relative path to packages folder")
}

// Execute run Gutenberg
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
