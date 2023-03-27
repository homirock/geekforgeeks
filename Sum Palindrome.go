// Given a number, reverse it and add it to itself unless it becomes a palindrome or number of iterations becomes more than 5.

// Example 1:

// Input: n = 23
// Output: 55 
// Explanation: reverse(23) = 32,then 32+23
// = 55 which is a palindrome. 
package main
import "fmt"

func sumPalindrom(num int){
	c:=0
	rev:= reverse(num)
	check:=rev+num
	check1:=reverse(check)
	if check==check1{
		fmt.Println("Sum is palindromic",check,check1)
	}else if c>5{
		fmt.Println("iteration exceeded above 5")

	}else{
		c=c+1
       sumPalindrom(check1)
	}
}

func reverse(num int) int{
	revers:=0
	for num > 0{
		remainder:=num%10
		revers= (revers*10)+remainder
		num=num/10
	}
	return revers
}

func main(){
	sumPalindrom(329)
	
}