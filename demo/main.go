package main

import (
	"fmt"

	"github.com/patrickbucher/consgo"
)

func main() {
	// Cells
	numbers := consgo.Cons(3, nil)
	numbers = consgo.Cons(2, numbers)
	numbers = consgo.Cons(1, numbers)
	fmt.Println(consgo.Slice(numbers))
	fmt.Println(consgo.Slice(consgo.List([]int{1, 2, 3, 4, 5, 6, 7})))

	// Streams
	integers := consgo.New(0, func(x int) int { return x + 1 })
	fmt.Println(consgo.StreamRef(integers, 10))
	fmt.Println(consgo.StreamTake(integers, 5))

	threes := consgo.New(3, func(x int) int { return x + 3 })
	fmt.Println(consgo.StreamTake(threes, 10))

	fibs := fibgen(1, 1)
	fmt.Println(consgo.StreamTake(fibs, 20))
}

func fibgen(a, b int) *consgo.Stream[int] {
	return consgo.ConsStream(a, func() *consgo.Stream[int] { return fibgen(b, a+b) })
}
