// Given an array of size N and you have to tell whether the array is perfect or not. An array is said to be perfect if it's reverse array matches the original array. If the array is perfect then print "PERFECT" else print "NOT PERFECT".

// Example 1:

// Input : Arr[] = {1, 2, 3, 2, 1}
// Output : PERFECT
// Explanation:
// Here we can see we have [1, 2, 3, 2, 1] 
// if we reverse it we can find [1, 2, 3, 2, 1]
// which is the same as before.
// So, the answer is PERFECT.

package main
import (
	"fmt"
)

// https://gosamples.dev/copy-slice/  (important)


func perfectArray(arr []int){
	arr1:=make([]int, len(arr))
	l:=len(arr)
	copy(arr1, arr)
	//arr3:=arr
	count:=0
	
	// fmt.Printf("array %[1]v,address %[1]p\n",arr1)
	// fmt.Printf("array %[1]v,address %[1]p\n",arr)
	// fmt.Printf("array %[1]v,address %[1]p\n",arr3)
	m:=(l-1)/2
	for i:=0;i<=m;i++{
		arr[i],arr[l-1-i]=arr[l-1-i],arr[i]
	}

	fmt.Println(arr)
	// fmt.Println(arr1)
	for k,v:= range arr{
		if v==arr1[k]{
			count++
			fmt.Println(count)
		}
		}
		if count ==(l){
			fmt.Println("PERFECT")
		}else{
			fmt.Println("NOT PERFECT")
		}
	}

func main(){
	perfectArray([]int{1, 2, 3, 2, 5})
	//time.Sleep(10000000)
}