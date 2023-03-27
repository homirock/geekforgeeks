package main
import(
	 "fmt"
	 "strconv"
)

func digit(num int){
	var count int
	for num>0{
		num=num/10
		count=count+1
	}
	fmt.Println(count)
}
//or using strconv
func digit1(num int){
	fmt.Println(len(strconv.Itoa(num)))
}

func main(){
	digit(12)
	digit1(12)
}
