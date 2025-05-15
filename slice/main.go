package main

import "fmt"

func main() {
	// s := []int{1, 2, 3} //slice literal
	// fmt.Println("slice", s, "len", len(s), "cap", cap(s))

	// s := make([]int, 3) //[0, 0, 0], len = 3, cap = 3
	// s[0] = 5            //[5, 0, 0], len = 3, cap = 3

	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s := make([]int, 3, 5) //[0, 0, 0], len = 3, cap = 5
	// s[0] = 5               //[5, 0, 0], len = 3, cap = 5
	// s[2] = 10              //[5, 0, 10], len = 3, cap = 5

	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	var s []int      //empty slice or nill slice [1]
	s = append(s, 1) //[1]
	fmt.Println(s)

}

//slice has total 3 element in it, 1. pointer, 2. length, 3.capacity

/*
1. slice from an existing array
2. slice from a slice
3. s := []int{1, 2, 3} //slice literal
4. if u use slice literal then length and capacity will be same
5. The pointer is to the first element of the slice or slices first element is the pointer.
6. The slice's pointer always points to its first element
7. Make func with length and capacity
8. In slice you can't set more capacity than the actual size of the array, if u do it then it will throw run time error
9. var s[] int		//empty slice or nill slice
*/

/*
1.Every function creates a stack frame
*/
