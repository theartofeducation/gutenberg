package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// ShellExec executes the given command in the OS shell.
func ShellExec(command string, args ...string) {
	shellCommand := exec.Command(command, args...)
	output, err := shellCommand.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}

// GetPackageFolders returns a list of folders that contain package.json files.
func GetPackageFolders(packagesFolder string) []os.FileInfo {
	files, err := ioutil.ReadDir("./" + packagesFolder)
	if err != nil {
		log.Fatal(err)
	}

	var directories []os.FileInfo

	for _, f := range files {
		if f.IsDir() {
			directories = append(directories, f)
		}
	}

	var packageFolders []os.FileInfo

	for _, pf := range directories {
		packageJSONPath := fmt.Sprintf("./%s/%s/package.json", packagesFolder, pf.Name())
		_, err := os.Stat(packageJSONPath)
		if !os.IsNotExist(err) {
			packageFolders = append(packageFolders, pf)
		}
	}

	return packageFolders
}
