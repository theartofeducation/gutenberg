package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Init initialize the core package with shared resources.
func Init() {
	log = logrus.New()
}

// ShellExec executes the given command in the OS shell.
func ShellExec(workingDir string, command string, args ...string) {
	var commandOutput bytes.Buffer
	var commandError bytes.Buffer
	shellCommand := exec.Command(command, args...)
	shellCommand.Dir = workingDir
	shellCommand.Stdout = &commandOutput
	shellCommand.Stderr = &commandError

	log.Infof("Executing command %s", shellCommand.String())
	err := shellCommand.Run()

	if err != nil {
		log.Errorf("ERROR executing command %q: %s (%s)\n", command, commandError.String(), err.Error())
	}

	log.Infof("%s\n", commandOutput.String())
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
