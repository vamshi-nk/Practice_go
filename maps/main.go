package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")
	// map literal

	teams := map[string][]string{
		"john" : []string{"mike","tyson","ali"},
		"lesner" :[]string{"broock","cena","reigns"},
	}
	fmt.Println(teams)
	// make keyword with size, but like slices it can grow past by adding key values
	// ages := make([string][]string ,10)  length 0 ,size 10

	// Reading and writing maps
	
	totalWins := map[string]int{}
	totalWins["camel"] = 1
	totalWins["zebra"] = 2
	fmt.Println(totalWins["camel"])
	fmt.Println(totalWins["zebra"])

	
}
