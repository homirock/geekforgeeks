// Given an array Arr[] of N integers.Find the sum of values of even and odd index positions separately.

// Example 1:

// Input:
// N=5
// Arr={1,2,3,4,5}
// Output:
// 9 6

package main 
import "fmt"

func EvenOddSum(arr []int)(int,int){
l:= len(arr)
sum1,sum2:=0,0
for i:=0;i<=l-1;i++{
	if arr[i]%2 == 0{
		sum1= sum1+arr[i]
	}else{
		sum2=sum2+arr[i]
	}
}
return sum1,sum2
}

func main(){
	arr := []int{1,2,4,5,9,8}
	x,y:=EvenOddSum(arr)
	fmt.Println(x,y)
}