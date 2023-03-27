package main
import "fmt"

func secondLargest(elements []int){

	max:=elements[0]
	//var ans int
	ans:=0
	fmt.Println(ans)
	for i:=0;i<len(elements);i++{
		if elements[i]>max{
			max = elements[i]
		}
	}
	for i:=0;i<len(elements);i++{
	if elements[i]>ans  && elements[i]!=max{
		ans = elements[i]
	}
	}
	fmt.Println(ans)
}

func main(){
	arr:= []int{1,5,3,2,4}
	secondLargest(arr)
}