package main

import "fmt"

func triplet(arr []int) bool {
	l := len(arr)
	for i := 0; i <= l-3; i++ {
		for j := i + 1; j <= l-2; j++ {
			for k := j + 1; k <= l-1; k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					return true
				}
			}
		}
	}
	return false

}

func main() {
	isPresent := triplet([]int{1, 2, 3, 4, 0, -1})
	fmt.Println(isPresent)
}