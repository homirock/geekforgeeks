package main

import (
	"fmt"
	"strings"
)

func main() {
	st := strings.Split("My Name is Partha", " ")
	fmt.Println(st)

	for i, j := 0, len(st)-1; j > i; i, j = i+1, j-1 {
		st[i], st[j] = st[j], st[i]
	}
	str1 := ""
	for _, v := range st {
		str1 = str1 + " " + v
	}

	fmt.Println(str1)
}