// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
// In this example, we shortened
package main

import "fmt"

func add(x, y, m int) int {
	return x + y + m
}

func main() {
	fmt.Println(add(42, 13, 12))
}
