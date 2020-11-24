package core

// RunAction will execute the given action (yarn task)
// in each of the package folders contained in the given packagesFolder.
func RunAction(action string) {
	actions := map[string]func(){
		"version": VersionPackage,
		"tag":     TagPackage,
	}

	if _, exists := actions[action]; !exists {
		log.Infof("The %s action is not yet implemented\n", action)
	}

	actions[action]()
}

// VersionPackage increments the version of the package in the given packageFolder.
func VersionPackage() {
	rootPackageJSONData := ReadPackageJSON(workingDirectory)

	YarnRunVersion(rootPackageJSONData)
}

// TagPackage adds a Git tag to a package with the updated version.
func TagPackage() {
	rootPackageJSONData := ReadPackageJSON(workingDirectory)

	log.Infof("Tagging package %s with v%s\n", rootPackageJSONData.Name, rootPackageJSONData.Version)

	// GitTag(packageFolder, rootPackageJSONData.Version)
}
