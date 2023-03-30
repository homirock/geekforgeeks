// Given a sorted array of size N and an integer K, find the position(0-based indexing) at which K is present in the array using binary search.

// Example 1:

// Input:
// N = 5
// arr[] = {1 2 3 4 5} 
// K = 4
// Output: 3
// Explanation: 4 appears at index 3.

// Example 2:

// Input:
// N = 5
// arr[] = {11 22 33 44 55} 
// K = 445
// Output: -1
// Explanation: 445 is not present.

package main

import "fmt"

func BinarySearch(arr []int,m int){
	count:=0
	for k,v:=range arr{
		if v==m{
			fmt.Println(k)
			count=count+1
			break
		}
	}
	if count==0{
		fmt.Println(-1)
	}

}
func main() {
	arr:=[]int{1,2,3,4,5,6,7,8,9}
	BinarySearch(arr,5)
}