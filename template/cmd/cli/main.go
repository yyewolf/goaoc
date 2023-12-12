package main

import (
	"aocli/template/internal/cli"
	"aocli/template/internal/markdown"
	"fmt"
)

func main() {

	fmt.Println(markdown.GenerateStars("2023"))
	cli.Execute()
}
