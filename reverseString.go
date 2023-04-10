// Golang program to reverse a string
package main

// importing fmt
import "fmt"

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) string{
	result:=""
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func main() {

	// Reversing the string.
	str := "Geeks"

	// returns the reversed string.
	 reverse:=reverse(str)
	 fmt.Println("string:",str)
	 fmt.Println("reverse:",reverse)
	
}
