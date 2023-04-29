package main

import (
	"fmt"
	"sort"
)

func removeDuplicate(arr []int)[]int{
	sort.Ints(arr)
	prev:=arr[0]
	arr1:=[]int{arr[0]}
	for i := 1; i < len(arr); i++{
		if arr[i]!=prev{
			arr1=append(arr1,arr[i])
			prev=arr[i]
		}

	}
	return arr1
}

func main() {

	arr := []int{2, 2, 2, 1, 1, 1, 3, 3, 3, 8, 8, 8}

	fmt.Println(removeDuplicate(arr))
	
}
