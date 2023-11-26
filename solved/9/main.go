package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Printf("isPalindrome(x, y): %v\n", isPalindrome(x))
}
