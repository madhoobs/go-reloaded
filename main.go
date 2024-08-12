package main

import (
	refine "go-reloaded/pkg"

	"os"
)

func main() {
	args := os.Args[1:]

	text, _ := os.ReadFile(args[0])
	
	refined := refine.AllChecks(string(text))

	handle := os.WriteFile(args[1], []byte(refined), 0644)
	if handle != nil {
		panic(handle)
	}
}
