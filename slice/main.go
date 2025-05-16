package main

import "fmt"

func print (numbers ...int){
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main(){
	print(1, 2, 3, 4, 5, 6)
}


//----------------------------------------------------------------------
// func main() {
// 	// s := []int{1, 2, 3} //slice literal
// 	// fmt.Println("slice", s, "len", len(s), "cap", cap(s))

// 	// s := make([]int, 3) //[0, 0, 0], len = 3, cap = 3
// 	// s[0] = 5            //[5, 0, 0], len = 3, cap = 3

// 	// fmt.Println(s)
// 	// fmt.Println(len(s))
// 	// fmt.Println(cap(s))

// 	// s := make([]int, 3, 5) //[0, 0, 0], len = 3, cap = 5
// 	// s[0] = 5               //[5, 0, 0], len = 3, cap = 5
// 	// s[2] = 10              //[5, 0, 10], len = 3, cap = 5

// 	// fmt.Println(len(s))
// 	// fmt.Println(cap(s))

// 	// var s []int      //empty slice or nill slice [1]
// 	// s = append(s, 1) //[1]
// 	// fmt.Println(s)
// }
//----------------------------------------------------------------------
// func main() {
// 	var x []int      //[], len = 0, cap = 0
// 	x = append(x, 1) //[1], len = 1, cap =1
// 	x = append(x, 2)
// 	x = append(x, 3)

// 	y := x

// 	x = append(x, 4)
// 	fmt.Println(y, len(y), cap(y))
// 	y = append(y, 5)

// 	x[0] = 10

// 	fmt.Println(x) //[1 2 3 5]
// 	fmt.Println(y) //[1 2 3 5]
// }
//----------------------------------------------------------------------
// func changeSlice(p []int) []int {
// 	p[0] = 10
// 	p = append(p, 11)
// 	return p

// }

// func main() {
// 	x := []int{1, 2, 3, 4, 5}
// 	x = append(x, 6)
// 	x = append(x, 7)

// 	a := x[4:]

// 	y := changeSlice(a)

// 	fmt.Println(x) //[1 2 3 4 10 6 7]
// 	fmt.Println(y) //[10 6 7 11]
// 	fmt.Println(x[0:8])
// }
//----------------------------------------------------------------------
/*
//slice has total 3 element in it, 1. pointer, 2. length, 3.capacity
1. slice from an existing array
2. slice from a slice
3. s := []int{1, 2, 3} //slice literal
4. if u use slice literal then length and capacity will be same
5. The pointer is to the first element of the slice or slices first element is the pointer.
6. The slice's pointer always points to its first element
7. Make func with length and capacity
8. In slice you can't set more capacity than the actual size of the array, if u do it then it will throw run time error
9. var s[] int		//empty slice or nill slice
10. Slice underlying array rule => 1024 -> 100% increase, 1024 < 25% increase
11.When you append an element to a slice in Go and its length equals its capacity, the Go runtime allocates a new underlying array with a larger capacity.
	I. For capacities up to 1024, the capacity typically doubles.
	II. For capacities above 1024, it grows by approximately 25% each time.
*/

/*
1.Every function creates a stack frame
*/
