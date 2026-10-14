package main

import (
	"fmt"
)

func main() {
	sum := 0
	ch := make(chan int)

	for i := 0; i < 1000; i++ {
		go func() {
			ch <- 1
		}()
	}

	for i := 0; i < 1000; i++ {
		sum += <-ch
	}

	fmt.Println(sum)
}
