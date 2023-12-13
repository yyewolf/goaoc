package main

import (
	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}
