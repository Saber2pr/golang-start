package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/favar/lambda"
)

func main() {
	str := `
	this is version:v0.0.1 branch.
	this is version:v0.0.2 branch.
	this is version:v0.0.3 branch.`

	reg := regexp.MustCompile(`version:v(\d+\.\d+\.\d+)`)
	if reg == nil {
		log.Fatal("regexp compile error")
	}

	isMatch := reg.MatchString(str)
	fmt.Println(isMatch)

	res1 := reg.FindAllString(str, 2)
	fmt.Println(res1)

	res2 := reg.FindAllString(str, -1)
	fmt.Println(res2)

	res3 := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(res3)

	res4 := reg.FindAllStringSubmatch(str, -1)
	res4Out := lambda.LambdaArray(res4).Map(func(el []string) string {
		return el[1]
	}).Pointer().([]string)
	fmt.Println(res4Out)
}
