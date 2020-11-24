package core

const yarnCommand = "yarn"

// YarnRun executes the "yarn run" command in the shell.
func YarnRun(workingDir string) {
	commandArgs := []string{"run"}

	ShellExec(yarnCommand, commandArgs...)
}

// YarnRunVersion executes the "yarn run version" command in the shell.
func YarnRunVersion(packageJSONData PackageJSON) {
	commandArgs := []string{}

	if len(packageJSONData.Workspaces) > 0 {
		commandArgs = append(commandArgs, "workspaces")
	}

	commandArgs = append(commandArgs, "run")
	commandArgs = append(commandArgs, "version")

	ShellExec(yarnCommand, commandArgs...)

	log.Infoln("Version(s) has been incremented, and CHANGELOG file(s) have been updated. Please review the changes that were made and commit them, and continue on to the next step in the process: `gutenberg tag`.")
}
