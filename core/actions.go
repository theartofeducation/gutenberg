package core

// RunTaskForPackages will execute the given action (yarn task)
// in each of the package folders contained in the given packagesFolder.
func RunTaskForPackages(packages []string, action string) {
	actions := map[string]func(string){
		"version": VersionPackage,
		"tag":     TagPackage,
	}

	for _, pf := range packages {
		if _, ok := actions[action]; !ok {
			log.Infof("The %s action is not yet implemented\n", action)
		}

		actions[action](pf)
	}
}

// VersionPackage increments the version of the package in the given packageFolder.
func VersionPackage(packageFolder string) {
	packageJSONData := ReadPackageJSON(packageFolder)

	log.Infof("Updating package %s\n", packageJSONData.Name)

	YarnRunVersion(packageFolder)
}

// TagPackage adds a Git tag to a package with the updated version.
func TagPackage(packageFolder string) {
	packageJSONData := ReadPackageJSON(packageFolder)

	log.Infof("Tagging package %s with v%s\n", packageJSONData.Name, packageJSONData.Version)

	GitTag(packageFolder, packageJSONData.Version)
}
