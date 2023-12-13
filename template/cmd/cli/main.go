package main

import (
	"aocli/template/internal/cli"
	"aocli/template/internal/folder"
	"os"
)

var root = folder.FindRoot()

func init() {
	// Change working directory to root
	err := os.Chdir(root)
	if err != nil {
		panic(err)
	}
}

func main() {
	// folder.CreateDay("2023", "01")
	cli.Execute()
}
