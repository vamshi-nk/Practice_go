package main

import(
	"fmt"
)

func main(){

	var a string = "Hello there"
	var b byte = a[6] //  prints ascii value  for byte 

	var a2 string = a[4:7]
	var a3 string = a[:5]
	var a4 string = a[6:]

	fmt.Println(
		"s :",a ,
		"\n b :",b, 
		"\n a1 :",a, 
		"\n a2 :",a2, 
		"\n a3 :",a3, 
		"\n a4 :",a4,
	)
	fmt.Printf("%c\n ",b) // prints t char
	
	// hello 🌞

	var c  string = "Hello 🌞"
	var c2 string = c[4:7]
	var c3 string = c[:5]
	var c4 string = c[6:]

	fmt.Println(c,c2,c3,c4)
	fmt.Println(len(c)) // prints 11 cause takes four bytes to represent the sun imoji

	// type conversion

	var d rune = 'x'
	var e string = string(a)
	var f byte = 'y'
	var g string = string(f)
	fmt.Printf(
    "%v %T\n%v %T\n%v %T\n%v %T\n",
    d, d,
    e, e,
    f, f,
    g, g,
	)
}	
