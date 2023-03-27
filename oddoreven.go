// Given a positive integer N, determine whether it is odd or even.
package main
import "fmt"
func oddeven(num int){
	if num%2==0{
		fmt.Printf("%d is even number",num)
	}else{
		fmt.Printf("%d is odd number",num)
	}
}

func main(){
	oddeven(87)
}