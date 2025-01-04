package main

import "fmt"

type Num int

const (
	_ Num = 1 << (10 * iota)
	KB
	MB
)

func main() {
	mask := 1 << 0

	result := KB | Num(mask) // setting 0. bit to 1

	fmt.Println(KB)
	fmt.Println(result)
}
