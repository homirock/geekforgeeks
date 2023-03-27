// Given a number N. The task is to complete the function convertFive() which replaces all zeros in the number with 5 and returns the number.

// Example:

// Input
// 1004
// 121

// Output
// 1554
// 121

package main
import "fmt"

func replaceDigit(num int){
//120
	remainder = num%10
	if remainder==0{
		num=num*10+5
	}
}