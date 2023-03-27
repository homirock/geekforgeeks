// Given an array Arr of N positive integers. Your task is to find the elements whose value is equal to that of its index value ( Consider 1-based indexing ).

// Note: There can be more than one element in the array which have the same value as its index. You need to include every such element's index. Follows 1-based indexing of the array.

// Example 1:

// Input:
// N = 5
// Arr[] = {15, 2, 45, 12, 7}
// Output: 2
// Explanation: Only Arr[2] = 2 exists here.

package main
import "fmt"

func indexedNumber(arr []int){
l:= len(arr)
	for i:=0;i<=l-1;i++{
		if arr[i]==i+1{
			fmt.Println(arr[i])
		}
	}
}

func main(){
	arr:= []int{3,2,7,4,8}
	indexedNumber(arr)
}
