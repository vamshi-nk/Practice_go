package main 

import(
	"fmt"
	"slices"
)

  func main(){
	var t = []int{12,32,23,}

	var u = []int {1:34,134,10:500}

	fmt.Println(t,u)

	// slice comparisions  == we use slices.Equal func

	x := []int {1,2,3,4,5}
	y := []int {1,2,3,4,5}
	z := []int {1,2,3,4,5,6}
	//s := []string{"a","b","c"}

	fmt.Println(slices.Equal(x,y)) // prints true
	fmt.Println(slices.Equal(x,z)) // prints false
	//fmt.Println(slices.Equal(x,s)) // doesn't compile


	

	// len
	//fmt.Println("Length of y :",len(y))

	//append
	var v []int
	v = append(v, 1)
	v = append(v, 2,3,4,5)

	fmt.Println(v)

	//len
	var s = []int{2,3,4,5,6,}
	lengthofs := len(s)
	fmt.Println("length of s =",lengthofs)

	// make keyword

	var a = make([]int ,5 ,10) // we can initlize the leand cap with make
	a = append([]int {12,14,16,18,20},a...)

	fmt.Println(a)

	// clear slice

	b := []string {"Sunday","Monday","Tuesday"}
	fmt.Println(b,len(b), cap(b))

	clear(b) // empty the slice
	fmt.Println(b, len(b), cap(b))

	var c  []int
	d:= []int {}

	fmt.Println(c,d)

	e := make([]int, 0 ,10)
	fmt.Println(e,len(e),cap(e))

	// sub slice
	// doen't make a copy instead the two variables are sharing the memory
	f := []string {"a", "b", "c", "d"}	// main sllice

	//sub slices
	g := f[:2]	// prints a b
	h := f[1:]	// prints b c d
	i := f[1:3]	// prints b c 
	j := f[:]	// prints a b c d

	fmt.Println("f :", f)
	fmt.Println("g :", g)
	fmt.Println("h :", h)
	fmt.Println("i :", i)
	fmt.Println("j :", j)
	
	// slices eith overlapping storage

	k := []string {"a","b","c","d"}
	l := k[:2]
	m := k[1:]

	k[1] = "y"
	l[0] = "x"
	m[1] = "z"

	fmt.Println("k : ",k)
	fmt.Println("l : ",l)
	fmt.Println("m : ",m)

	/* output
	k :  [x y z d]
	l :  [x y]
	m :  [y z d]
	*/

	// append makes overlaping slices more confusing

	n := []string {"a","b","c","d"}
	
	o := n[:2]
	fmt.Println(cap(n),cap(o)) 

	o = append(o, "z")		// appends after first two elements
	fmt.Println("n:",n)		// prints abzd cause append ,because subslice shares same memory
	fmt.Println("o:",o)
	/*
		4 4
		n: [a b z d]
		o: [a b z]
	*/

	//even more confusing slices

	p := make([]string ,0,5)
	p = append(p, "q","r","s","t")

	
}