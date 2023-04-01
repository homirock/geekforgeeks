package main
import "fmt"
func missingNum(arr []int,n int){
	Acutal_sum:=(n*(n+1))/2
	Array_sum:=0
	for i := 0; i<len(arr) ; i++{
		Array_sum=Array_sum+arr[i]
	}
	missing_num:=Acutal_sum-Array_sum
	fmt.Println("missing number is:",missing_num)
}

func main(){
	missingNum([]int{1,2,3,4,6},6)
}