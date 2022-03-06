package main

import (
	"fmt"
	"golangstart/pkg/argv"
	"golangstart/pkg/json"
	"log"
)

func main() {
	args := argv.GetArgs()
	if len(args) < 2 {
		log.Fatal("need input")
	}

	var path = args[1]

	jsonRes1, err := json.GetJsonValue(path, "name")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonRes1.MustString())

	jsonRes2, err := json.GetJsonValue(path, "scripts", "start")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonRes2.MustString())
}
