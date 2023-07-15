package main

import (
	"fmt"

	"github.com/patrickbucher/consgo"
)

func main() {
	numbers := consgo.Cons(3, nil)
	numbers = consgo.Cons(2, numbers)
	numbers = consgo.Cons(1, numbers)
	fmt.Println(consgo.Slice(numbers))
	fmt.Println(consgo.Slice(consgo.List([]int{1, 2, 3, 4, 5, 6, 7})))
}
