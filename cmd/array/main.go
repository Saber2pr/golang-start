package main

import (
	"fmt"
	"strconv"

	"github.com/favar/lambda"
)

func main() {
	array := []int{1, 5, 2, 4}
	larr := lambda.LambdaArray(array)

	// map
	res1 := larr.Map(func(el int) string {
		return "val:" + strconv.Itoa(el)
	}).Pointer().([]string)

	fmt.Printf("%v\n", res1)

	// filter
	res2 := larr.Filter(func(el int) bool {
		return el > 3
	}).Pointer().([]int)

	fmt.Printf("%v\n", res2)

	// reduce
	res3 := larr.Sum(func(el int) int {
		return el
	}).(int)

	fmt.Printf("%v\n", res3)

	// sort
	res4 := larr.Sort(func(a int, b int) bool {
		return a < b
	}).Pointer().([]int)

	fmt.Printf("%v\n", res4)
}
