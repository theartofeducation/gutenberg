package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// PackageJSON package.json data
type PackageJSON struct {
	Name    string
	Version string
}

// ReadPackageJSON reads the `name` and `version` values from the package.json file
// at the given filepath, and returns a packageJSON struct with those values
func ReadPackageJSON(filepath string) PackageJSON {
	var packageJSONData PackageJSON

	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err = json.Unmarshal(fileData, &packageJSONData); err != nil {
		fmt.Println(err.Error())
	}

	return packageJSONData
}
