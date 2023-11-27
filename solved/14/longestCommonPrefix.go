package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	t := strs[0][:0]
	var min int = len(strs[0])
	for i := range strs {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	for i := 1; i <= min; i++ {
		t = strs[0][:i]
		fmt.Printf("t: %v\n", t)
		for m := range strs {
			if strs[m][:i] != t {
				return t[:i-1]
			}
		}
	}
	return t
}
