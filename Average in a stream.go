// Given a stream of incoming numbers, find average or mean of the stream at every point.

 

// Example 1:

// Input:
// n = 5
// arr[] = {10, 20, 30, 40, 50}
// Output: 10.00 15.00 20.00 25.00 30.00 
// Explanation: 
// 10 / 1 = 10.00
// (10 + 20) / 2 = 15.00
// (10 + 20 + 30) / 3 = 20.00
// And so on.
package main
import "fmt"

var sum int
var n int

func getAvg(val,m int) int{
	
	sum=sum+val
	return sum/m
}

func average(arr []int){
	
	arr1:=[]int{}
	l:= len(arr)
	for i:=0;i<=l-1;i++{
		avg:=getAvg(arr[i],i+1)
		arr1=append(arr1,avg)
	}
	fmt.Println(arr1)
}

func main(){
	arr:=[]int{10,20,30,40,50}
	average(arr)
}