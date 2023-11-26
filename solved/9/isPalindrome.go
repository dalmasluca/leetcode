package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var tmp int = x
	var y int
	for tmp != 0 {
		y *= 10
		y += (tmp % 10)
		tmp /= 10
	}
	return x == y
}
