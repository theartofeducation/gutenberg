package core

var gitCommand string = "git"

// GitLog runs "git log" in the shell.
func GitLog(workingDir string) {
	commandArgs := []string{"log"}

	ShellExec(workingDir, gitCommand, commandArgs...)
}

// GitLogShort runs "git log --pretty=oneline --abbrev-commit" in the shell.
func GitLogShort(workingDir string) {
	commandArgs := []string{
		"log",
		"--pretty=oneline",
		"--abbrev-commit",
	}

	ShellExec(workingDir, gitCommand, commandArgs...)
}

// GitTag executes "git tag" for the specified version.
func GitTag(workingDir, version string) {
	commandArgs := []string{
		"tag",
		"-a",
		"v" + version,
		"-m",
		`"Release v`+version+`"`,
		"--sign",
	}

	ShellExec(workingDir, gitCommand, commandArgs ...)
}
