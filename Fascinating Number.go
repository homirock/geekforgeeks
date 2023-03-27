// Given a number N. Your task is to check whether it is fascinating or not.
// Fascinating Number: When a number(should contain 3 digits or more) is multiplied by 2 and 3, and when both these products are concatenated with the original number, then it results in all digits from 1 to 9 present exactly once.

// Example 1:

// Input: 
// N = 192
// Output: Fascinating
// Explanation: After multiplication with 2
// and 3, and concatenating with original
// number, number will become 192384576 
// which contains all digits from 1 to 9.

package main
import (
	"fmt"
	"strconv"
)

func fascinating(num int){
	num1:= strconv.Itoa(num)
	num2:= strconv.Itoa(num*2)
	num3:= strconv.Itoa(num*3)
	finalstringnum:=num1+num2+num3
	var sum int
	fmt.Println(finalstringnum)
	//finalintnum:=strconv.Atoa(finalstringnum)
	for _,v:=range finalstringnum{
		fmt.Println(string(v))
		element,_:=strconv.Atoi(string(v))
		sum=sum+element
	}
	if sum==45{
		fmt.Println("Fascinating")
	}else{
		fmt.Println("Not Fascinating")
	}

}

func main(){
	fascinating(192)
}