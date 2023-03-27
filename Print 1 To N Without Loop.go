// Print numbers from 1 to N without the help of loops.

// Example 1:

// Input:
// N = 10
// Output: 1 2 3 4 5 6 7 8 9 10

package main
import "fmt"

func printNum(num int){
	
	if num>0{
		printNum(num-1) //recursion lIFO
		fmt.Println(num)
		
	}
}

func main(){
	printNum(10)
}