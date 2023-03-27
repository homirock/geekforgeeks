// Write a program to find the sum of the given series 1+2+3+ . . . . . .(N terms) 

// Example 1:

// Input:
// N = 1
// Output: 1
// Explanation: For n = 1, sum will be 1.
// Example 2:

// Input:
// N = 5
// Output: 15
// Explanation: For n = 5, sum will be 1
// + 2 + 3 + 4 + 5 = 15.

package main
import "fmt"

func seriesSum(num int){
	sum:=0
	for i:=1; i<=num ;i++{
		sum = sum+i
	}
	fmt.Println("sum:",sum)

}

func main(){
	seriesSum(5)
}