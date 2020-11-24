package core

import (
	"fmt"
	"path"
)

// RunTaskForPackages will execute the given action (yarn task)
// in each of the package folders contained in the given packagesFolder.
func RunTaskForPackages(packages []string, action string) {
	actions := map[string]func(string){
		"version": VersionPackage,
	}

	for _, pf := range packages {
		if _, ok := actions[action]; !ok {
			fmt.Printf("The %s action is not yet implemented\n", action)
		}

		actions[action](pf)
	}
}

// VersionPackage increments the version of the package in the given packageFolder.
func VersionPackage(packageFolder string) {
	packageJSONFilePath := path.Join(packageFolder, "package.json")
	packageJSONData := ReadPackageJSON(packageJSONFilePath)

	fmt.Printf("Updating package %s\n", packageJSONData.Name)

	YarnRunVersion(packageFolder)
}
