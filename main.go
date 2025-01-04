package main

import "fmt"

type Num bool

func main() {
	num := 10
	num2 := 12

	mask := 1 << 2

	if num&mask != 0 {
		fmt.Println("3. bit is 1")
	} else {
		fmt.Println("3. bit is 0")
	}

	if num2&mask != 0 {
		fmt.Println("3. bit is 1")
	} else {
		fmt.Println("3. bit is 0")
	}
}
