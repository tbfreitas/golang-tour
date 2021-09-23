// A function can take zero or more arguments.

// In this example, add takes two parameters of type int.

// Notice that the type comes after the variable name.

// (For more about why types look the way they do, see the article on Go's declaration syntax.)
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mult(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(10, 9))
	fmt.Println(mult(3, 10))
}
