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
}
