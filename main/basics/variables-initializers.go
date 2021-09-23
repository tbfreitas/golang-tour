// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	var m = 1000

	fmt.Println(i, j, c, python, java, m)
}
