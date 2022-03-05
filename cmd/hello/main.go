package main

import (
	"fmt"
	"golangstart/pkg/hello"

	"github.com/hashicorp/go-uuid"
)

func main() {
	str := hello.Log()
	uuid, _ := uuid.GenerateUUID()
	fmt.Println(str + uuid)
}
