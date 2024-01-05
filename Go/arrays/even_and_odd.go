// For the given array of integers, count even and odd elements.

// Examples:

// Input:
// int arr[5] = {2, 3, 4, 5, 6}
// Output:
// Number of even elements = 3
// Number of odd elements = 2

// Input:
// int arr[5] = {22, 32, 42, 52, 62}
// Output:
// Number of even elements = 5
// Number of odd elements = 0

package main

import "fmt"

func main() {
	fmt.Println("Count number of even and odd elements in an array")

	arr := [5]int{2, 3, 4, 5, 6}

	var odd_count int = 0
	var even_count int = 0

	for i := range arr {
		if arr[i]&1 == 1 {
			odd_count += 1
		} else {
			even_count += 1
		}

	}
	fmt.Println("Number of even elements = ", even_count)
	fmt.Println("Number of odd elements = ", odd_count)

}
