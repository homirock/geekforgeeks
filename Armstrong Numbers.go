// For a given 3 digit number, find whether it is armstrong number or not. An Armstrong number of three digits is an integer such that the sum of the cubes of its digits is equal to the number itself. Return "Yes" if it is a armstrong number else return "No".
// NOTE: 371 is an Armstrong number since 33 + 73 + 13 = 371

package main
import (
	"fmt"
)

func Armstrong(num int) {
	input_num:= num
	result:=0
	for num > 0{
		remainder:=num%10
		result=result+(remainder*remainder*remainder)// 2^7
		num=num/10
	}

	if result == input_num{
		fmt.Println("Armstrong number")
	}else{
		fmt.Println("Not a Armstrong number")
	}
}

func main(){
	Armstrong(153)
}