package main

import(
	"fmt"
)

func palindrom(num int) bool{
	input_num:= num
	result:=0
	for num > 0{
		remainder:=num%10
		result= (result*10)+remainder
		num=num/10
	}

	if result==input_num{
		return true
	}else{
		return false
	}
}

func main(){
	status:=palindrom(121)
	fmt.Println(status)
}