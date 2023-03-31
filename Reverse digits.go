// Given N,  reverse the digits of N.
 

// Example 1:

// Input: 200
// Output: 2
// Explanation: By reversing the digts of 
// number, number will change into 2.
// Example 2:

// Input : 122
// Output: 221
// Explanation: By reversing the digits of 
// number, number will change into 221.
package main
import "fmt"

func reverseDigit(num int){
 sum:=0
	for num>0{
		remainder:=num%10
		sum=sum*10+remainder
		num=num/10
	}
	fmt.Println("reversed num:",sum)

}

func main() {
	reverseDigit(121)
}