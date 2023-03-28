// Given a list of names, display the longest name.


// Example:

// Input:
// N = 5
// names[] = { "Geek", "Geeks", "Geeksfor",
//   "GeeksforGeek", "GeeksforGeeks" }

// Output:
// GeeksforGeeks
package main
import "fmt"

func longestName(arr []string) {
	l:= len(arr)
	longest:=""
	for i := 0; i < l-1; i++ {
		for j := i+1; j <= l-1; j++ {
		if len(arr[j]) > len(arr[i]) {
			longest = arr[j]
		}
	}
}
fmt.Println(longest)
}

func main() {
	arr:=[]string{"Geek", "Geeks", "Geeksfor","GeeksforGeek", "GeeksforGeeks" }
	longestName(arr)
}