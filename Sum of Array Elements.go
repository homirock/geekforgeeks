// Given an integer array arr of size n, you need to sum the elements of arr.

// Example 1:

// Input:
// n = 3
// arr[] = {3 2 1}
// Output: 6

package main
import "fmt"

func sumArray(arr []int){
	l:= len(arr)
	sum:=0
	for i:=0;i<=l-1;i++{
		sum=sum+arr[i]
	}
	fmt.Println(sum)
}

func main(){
	sumArray([]int{1,3,7,8})
}