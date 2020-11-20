package tools

import "github.com/theartofeducation/gutenberg/core"

var gitCommand string = "git"

// GitLog runs "git log" in the shell
func GitLog() {
	commandArgs := []string{"log"}

	core.ShellExec(gitCommand, commandArgs...)
}

// GitLogShort runs "git log --pretty=oneline --abbrev-commit" in the shell
func GitLogShort() {
	commandArgs := []string{"log", "--pretty=oneline", "--abbrev-commit"}

	core.ShellExec(gitCommand, commandArgs...)
}
