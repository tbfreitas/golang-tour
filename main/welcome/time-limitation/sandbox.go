package main

import (
	"fmt"
	"time"
)

// There are also limits on execution time and on CPU and memory usage, and the program cannot access external network hosts.
func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}
