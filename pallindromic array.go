// retun 1 if all the elements of an array are palindromic

package main
import "fmt"

func palindromic(arr []int) int{
	l:=len(arr)
	var n int
	
	for i:=0;i<=l-1;i++{
		num:=arr[i]
		val:=0
		for num>0{
		reminder:=num%10
	   	val=val*10+reminder
	   	num= num/10
		}
		if val == arr[i]{	
			n++
			//fmt.Println(n)
		}
	}
	if n==l{
		return 1
	}else{
		return 0
	}
}

func main(){
	arr:=[]int{111,121,252,323,1881}
	status:= palindromic(arr)
	fmt.Println(status)
}