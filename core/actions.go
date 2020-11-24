package core

import (
	"fmt"
	"path"
	"time"

	"github.com/briandowns/spinner"
)

// RunTaskForPackages will execute the given action (yarn task)
// in each of the package folders contained in the given packagesFolder
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

// VersionPackage increments the version of the package in the given packageFolder
func VersionPackage(packageFolder string) {
	s := spinner.New(spinner.CharSets[14], 75*time.Millisecond)
	s.Suffix = " Getting package data"
	s.Start()

	packageJSONFilePath := path.Join(packageFolder, "package.json")
	packageJSONData := ReadPackageJSON(packageJSONFilePath)

	_ = s.Color("blue")
	s.Suffix = fmt.Sprintf(" Updating package %s", packageJSONData.Name)

	YarnRunVersion()

	s.FinalMSG = "✅ Package(s) updated"
	s.Stop()
}
