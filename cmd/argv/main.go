package main

import (
	"fmt"
	"golangstart/pkg/argv"
)

func main() {
	args := argv.GetArgs()
	fmt.Printf("%v", args)
}
