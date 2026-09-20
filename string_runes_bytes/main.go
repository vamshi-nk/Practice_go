package main

import(
	"fmt"
)

func main(){

	var s string = "Hello there"
	var b byte = s[6] //  prints ascii value  for byte 

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]

	fmt.Println(
		"s :",s ,
		"\n b :",b, 
		"\n s1 :",s, 
		"\n s2 :",s2, 
		"\n s3 :",s3, 
		"\n s4 :",s4,
	)
	fmt.Printf("%c ",b) // prints t char
	
}	
