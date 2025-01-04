package main

import "fmt"

type Num int

const (
	_ Num = 1 << (10 * iota)
	KB
	MB
)

func main() {
	mask := 3

	result := KB ^ Num(mask) // Toggling the 0. and 1. bit

	fmt.Println(KB)
	fmt.Println(result)

	mask2 := 1 << 10

	result = result ^ Num(mask2) // Toggling the 10. bit

	fmt.Println(result)
}
