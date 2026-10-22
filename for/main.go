package main

import (
	"fmt"
)

func main() {
	// while loop
	i := 1
	for i <= 11 {
		fmt.Println("Hello")
		i++
	}
	// for loop
	for i := 1; i < 11; i++ {
		fmt.Println("Hi")
	}
	// infinte loop
	// for {
	// 	fmt.Println("I love you")
	// }

	//range
	for range 5 {
		fmt.Println("I love you")
	}
}
