package core

const gitCommand = "git"

// GitLog runs "git log" in the shell.
func GitLog() {
	commandArgs := []string{"log"}

	ShellExec(gitCommand, commandArgs...)
}

// GitLogShort runs "git log --pretty=oneline --abbrev-commit" in the shell.
func GitLogShort() {
	commandArgs := []string{
		"log",
		"--pretty=oneline",
		"--abbrev-commit",
	}

	ShellExec(gitCommand, commandArgs...)
}

// GitTag executes "git tag" for the specified version.
func GitTag(version string) {
	commandArgs := []string{
		"tag",
		"-a",
		"v" + version,
		"-m",
		`"Release v` + version + `"`,
		"--sign",
	}

	ShellExec(gitCommand, commandArgs...)
}
