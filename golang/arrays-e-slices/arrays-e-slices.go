package main

import "fmt"

func main() {

	fmt.Println("Arrays e Slices")

	// ARRAY

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"teste"}
	fmt.Println(array2)

	// SLICE

	slice := []int{1}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	// ARRRAYS INTERNOS

	fmt.Println("---------------------------------")
	slice2 := make([]int, 10, 15)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // lenght
	fmt.Println(cap(slice2)) // capacity

	slice2 = append(slice2, 5)
	fmt.Println(slice2)

}
