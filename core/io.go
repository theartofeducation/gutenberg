package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const filename = "package.json"

// PackageJSON package.json data.
type PackageJSON struct {
	Name       string
	Version    string
	Workspaces []string
}

// ReadPackageJSON reads the `name` and `version` values from the package.json file
// at the given folder, and returns a PackageJSON struct with those values.
func ReadPackageJSON(packageFolder string) PackageJSON {
	filepath := path.Join(packageFolder, filename)
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var packageJSONData PackageJSON
	if err = json.Unmarshal(fileData, &packageJSONData); err != nil {
		fmt.Println(err.Error())
	}

	return packageJSONData
}
