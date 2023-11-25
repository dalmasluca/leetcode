package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	for i := 0; i < 2; i++ {
		if len(str1) > len(str2) {
			str1, str2 = str2, str1
		}
		if str2[:len(str1)] != str1 {
			return ""
		}
		if str1 == str2 {
			fmt.Printf("\"entro\": %v\n", "entro")
			return str1
		}
		str2 = str2[len(str1):]
	}
	return ""
}

func main() {
	fmt.Printf("%s", gcdOfStrings("A", "A"))
}
