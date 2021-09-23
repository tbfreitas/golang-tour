// Multiple results
// A function can return any number of results.

// The swap function returns two strings.
package main

import "fmt"

func swap(x, y, m string) (string, string, string) {
	return y, x, m
}

func main() {
	a, b, c := swap("hello", "world", "test")
	fmt.Println(a, b, c)
}
