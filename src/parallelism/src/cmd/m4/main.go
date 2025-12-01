package main

import (
	"fmt"
)

func main() {
	before()
	// after()
}

func before() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func after() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
