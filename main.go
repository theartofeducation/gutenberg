package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/theartofeducation/gutenberg/core"
)

func main() {
	log := logrus.New()
	workingDirectory, _ := os.Getwd()
	core.Execute(log, workingDirectory)
}
