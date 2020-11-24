package core

var gitCommand string = "git"

// GitLog runs "git log" in the shell.
func GitLog() {
	commandArgs := []string{"log"}

	ShellExec(gitCommand, commandArgs...)
}

// GitLogShort runs "git log --pretty=oneline --abbrev-commit" in the shell.
func GitLogShort() {
	commandArgs := []string{"log", "--pretty=oneline", "--abbrev-commit"}

	ShellExec(gitCommand, commandArgs...)
}
