// A and B are good friend and programmers. They are always in a healthy competition with each other. They decide to judge the best among them by competing. They do so by comparing their three skills as per their values. Please help them doing so as they are busy.

// Set for A are like [a1, a2, a3]
// Set for B are like [b1, b2, b3]

// Compare ith skill of A with ith skill of B
// if a[i] > b[i] , A's score increases by 1
// if a[i] < b[i] , B's score increases by 1
// if a[i] = b[i] , nothing happens
// Example 1:
// Input : 
// A = {4, 2, 7}
// B = {5, 6, 3}
// Output : 
// 1 2
// Example 2:
// Input : 
// A = {4, 2, 7}
// B = {5, 2, 8}
// Output : 
// 0 2

package main
import "fmt"

func competeSkills(arr1,arr2 []int)(int,int){
	a1,a2:=0,0

	l:=len(arr1)
	for i:=0;i<=l-1;i++{
		if arr1[i]>arr2[i]{
			a1=a1+1
		}else if arr1[i]<arr2[i] {
			a2=a2+1
		}
	}
	return a1,a2
}

func main(){

	arr1:= []int{4, 2, 7}
	arr2:= []int{2,2,4}
	x,y:= competeSkills(arr1,arr2)
	fmt.Println(x,y)
}