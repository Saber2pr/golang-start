package main

import (
	"fmt"
	"golangstart/pkg/argv"
	"golangstart/pkg/fs"
	"log"
)

func main() {
	args := argv.GetArgs()
	if len(args) < 2 {
		log.Fatal("need input")
	}

	var path = args[1]

	text, err := fs.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", text)
}
