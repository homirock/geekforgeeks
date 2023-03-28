// Given an sorted array A of size N. Find number of elements which are less than or equal to given element X.

// Example 1:

// Input:
// N = 6
// A[] = {1, 2, 4, 5, 8, 10}
// X = 9
// Output:
// 5.
package main

import "fmt"

func countSmaller(arr []int, num int) int {
	l := len(arr)
	count := 0
	for i := 0; i <= l-1; i++ {
		if arr[i] <= num {
			count = count + 1
		}
	}
	return count
}

func main() {
arr:= []int{1,2,3,4,5}
c:=countSmaller(arr,3)
fmt.Println(c)
}
