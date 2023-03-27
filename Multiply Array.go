// Calculate the product of all the elements in an array.

// Example 1:

// Input:
// 5
// 1 2 3 4 5
// Output:
// 120

package main
import "fmt"

func mul(arr []int) int{
	l:= len(arr)
	mul:=1
	for i:=0;i<=l-1;i++{
		mul=mul*arr[i]
	}
	return mul
}

func main(){
	arr:= []int{1,2,3,4,5}
	x:= mul(arr)
	fmt.Println(x)
}