package main
import "fmt"

func alterElement(arr []int){
	l:=len(arr)
	// for i:=0;i<=l-1;i=i+2{
	// 	fmt.Println(arr[i])
	// } or 
	for i:=0;i<=l-1;i++{
		if i%2==0{
		fmt.Println(arr[i])
		}
	}
}

func main(){
	arr:=[]int{1,2,3,4,5,7}
	alterElement(arr)
}