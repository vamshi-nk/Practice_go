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

	//q := p[:2]
	//r := p[2:]

	//modified

	q := p[:2:2]
	r := p[2:4:4]

	fmt.Println(cap(p),cap(q),cap(r))
	q = append(q, "i","j","k","l")
	p = append(p, "x")
	r = append(r, "y")

	fmt.Println("p : ",p)
	fmt.Println("q : ",q)
	fmt.Println("r : ",r)

	// before modify code output			// after output
	/*
		5 5 3								5 2 2
		p :  [q r s t y]					p :  [q r s t x]
		q :  [q r i j k l]					q :  [q r i j k l]
		r :  [s t y	]						r :  [s t y]
											
	*/

	// copy built-in function 

	ab := []int {1,2,3,4,5}
	cd := make([]int,4)
	num := copy(ab,cd)

	fmt.Println(y,num)

	// sub slice copy

	ef := []int {1,2,3,4}
	gh := make([]int ,2)
	nums := copy(ef,gh)
	
	fmt.Println("nums :",nums)

	// copy from middle

	ij := []int {1,2,3,4}
	kl := make([]int,2)
	numb:= copy(kl,ij[:2])

	fmt.Println(numb)

	//copied values print

	qw := []int{1,2,3,4}	// output
	er := [4]int{5,6,7,8}	//[5 6]
	ty := make([]int,2)		//[1 2 3 4]
	copy(ty,er[:])
	fmt.Println(ty)
	copy(er[:],qw)
	fmt.Println(er)

	// prac

	ui := []int {1,2,3,4}
	pa := make([]int,2)
	copy(pa,ui)
	fmt.Println(ui)
	fmt.Println(pa)

	// output
	// [1 2 3 4]
	// [1 2]

	// convertiong Array to slices

	xarray := [4]int{5,6,7,8}
	xslice := xarray[:]
	fmt.Println(xslice)

	// subset of array to slice

	as := [4]int{1,2,3,4}
	df := as[:2]
	fd := as[2:] 
	as[0] = 10
	fmt.Println("as :",as)
	fmt.Println("df :",df)
	fmt.Println("fd :",fd)
	
	// converting Slices to Arrays

	 xSlice := []int{1,2,3,4,}
	 xArray := [4]int(xSlice)

	 smallArray := [2]int(xSlice)
	 fullArray := [4]int(xSlice)
	 xSlice[0] = 10
	 
	 fmt.Println("xSlice :",xSlice)
	 fmt.Println("xArray :",xArray)
	 fmt.Println("smallArray :",smallArray)
	 fmt.Println("fullArray :",fullArray)

	 // runtime error if specify array len > slice
	 panicArray := [5]int(xSlice)
	 fmt.Println("panicArray :",panicArray)

	 // ouputs : 
	 // panic: runtime error: cannot convert slice with length 4 to array or pointer to array with length 5
  
	  
	
}
