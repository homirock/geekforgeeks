package main

import "fmt"

func main() {

	a := []string{"abc", "bcd", "cde", "def"}
	fmt.Println(a)
	l := len(a)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)

}