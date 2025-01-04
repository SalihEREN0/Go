package main

import "fmt"

type Num int

const (
	_ Num = 1 << (10 * iota)
	KB
	MB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
}
