//without using 3rd variable
package main
import "fmt"

func swapping(num1,num2 int)(int,int){

	num2 = num1-num2
	num1 = num1-num2
	num2 = num1+num2
	return num1,num2
}

func main(){
	a,b:=12,13
	fmt.Printf("the two numbers are %d and %d\n",a,b)
	x,y:= swapping(a,b)
	fmt.Printf("the two numbers are %d and %d",x,y)
}