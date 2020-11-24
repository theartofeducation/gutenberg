package core

var yarnCommand string = "yarn"

// YarnRun executes the "yarn run" command in the shell.
func YarnRun(workingDir string) {
	commandArgs := []string{"run"}

	ShellExec(workingDir, yarnCommand, commandArgs...)
}

// YarnRunVersion executes the "yarn run version" command in the shell.
func YarnRunVersion(workingDir string) {
	commandArgs := []string{"run", "version"}

	ShellExec(workingDir, yarnCommand, commandArgs...)

	log.Infoln("Version(s) has been incremented, and CHANGELOG file(s) have been updated. Please review the changes that were made and commit them, and continue on to the next step in the process: `gutenberg tag`.")
}
