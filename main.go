package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("Number from first function: ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("Number from second function: ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}
