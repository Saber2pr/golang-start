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

	err := json.SetJsonValue(path, "yarn build", "scripts", "build")
	if err != nil {
		log.Fatal(err)
	}

	jsonRes1, err := json.GetJsonValue(path, "scripts", "build")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonRes1.MustString())
}
