package main

import (
	"fmt"
	"golangstart/pkg/argv"
	"golangstart/pkg/fs"
	"log"
	"path/filepath"
)

func main() {
	args := argv.GetArgs()
	if len(args) < 2 {
		log.Fatal("need input")
	}

	var path = args[1]

	files, err := fs.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	filePaths := []string{}
	for i := 0; i < len(files); i++ {
		absFilePath, err := filepath.Abs(filepath.Join(path, files[i]))
		if err != nil {
			log.Fatal(err)
		}
		filePaths = append(filePaths, absFilePath)
	}

	fmt.Printf("%v", filePaths)
}
