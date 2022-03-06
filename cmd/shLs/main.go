package main

import (
	"fmt"
	"golangstart/pkg/shell"
	"log"
)

func main() {
	out1, err1 := shell.ExecCommand("go version")
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(out1)

	out2, err2 := shell.Ls("./")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(out2)
}
