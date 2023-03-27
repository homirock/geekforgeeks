// Given an array of N distinct elements, the task is to find all elements in array except two greatest elements in sorted order.
// Example 1:
// Input : 
// a[] = {2, 8, 7, 1, 5}
// Output :
// 1 2 5 
// Explanation :
// The output three elements have two or
// more greater elements.

package main
import "fmt"

func findElements(arr []int) []int{
	
	l:=len(arr)
	//s:=make(arr,l)
	fmt.Println(l)
	for i:=0;i<l-1;i++{ //0-->len-1
		for j:=i+1; j<=l-1;j++{ // 1--->l
			if arr[i]> arr[j]{
				arr[i],arr[j]=arr[j],arr[i]
			}
		}	
	}
	fmt.Println(arr)
	return arr[:l-2]
}

func main(){
	arr := []int{2,1,7,5,3}
	a:=findElements(arr)
	fmt.Println(a)
}