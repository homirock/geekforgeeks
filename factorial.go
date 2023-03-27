package main
import "fmt"

func factorial(num int) int{
fact:=1

for num>0{
	fact=fact*num
	num=num-1
}
return fact

}

func main(){
	x:=factorial(6)
	fmt.Println(x)
}