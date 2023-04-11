package main

import "fmt"

func reverseString(v string) string {
	result := ""
	for _, value := range v {
		result = string(value) + result
	}
	return result
}

func main() {

	a := []string{"abc", "bcd", "cde", "def"}
	var b []string
	fmt.Println(len(b))
	for _, v := range a {
		rev := reverseString(v)
		b = append(b, rev)
	}

	fmt.Println(b)
	l := len(b)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	fmt.Println(b)

}
