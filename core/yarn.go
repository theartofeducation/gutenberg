package core

var yarnCommand string = "yarn"

// YarnRun executes the "yarn run" command in the shell
func YarnRun() {
	commandArgs := []string{"run"}

	ShellExec(yarnCommand, commandArgs...)
}

// YarnRunVersion executes the "yarn run version" command in the shell
func YarnRunVersion() {
	commandArgs := []string{"run", "version"}

	ShellExec(yarnCommand, commandArgs...)
}
