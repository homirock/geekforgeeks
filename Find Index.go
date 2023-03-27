// Given an unsorted array Arr[] of N integers and a Key which is present in this array. You need to write a program to find the start index( index where the element is first found from left in the array ) and end index( index where the element is first found from right in the array ).If the key does not exist in the array then return -1 for both start and end index in this case.

// Example 1:

// Input:
// N = 6
// arr[] = { 1, 2, 3, 4, 5, 5 }
// Key = 5
// Output:  4 5
// Explanation:
// 5 appears first time at index 4 and
// appears last time at index 5.
// (0 based indexing)
 

// Example 2:

// Input:
// N=6
// arr[] = { 6, 5, 4, 3, 1, 2 }
// Key = 4
// Output:  2 2 

package main
import "fmt"

func findIndex(arr []int, key int)(int,int){
	index:=[]int{}
	l:=len(arr)
	for i:=0;i<=l-1;i++{
		if arr[i]==key{
			index=append(index,i)
		}
	}
	if len(index)==1{
		return index[0],index[0]
	}else{
		return index[0],index[len(index)-1]
	}
}

func main(){
	arr:=[]int{ 6, 5, 4, 3, 1, 3 }
	x,y:=findIndex(arr,3)
	fmt.Println(x,y)
}