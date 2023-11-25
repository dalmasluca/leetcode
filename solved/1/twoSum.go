package main

import "fmt"

type record struct {
	found bool
	idx   int
}

func twoSum(nums []int, target int) []int {
	var m map[int]record
	var i int
	m = make(map[int]record, len(nums))

	for i = range nums {
		m[nums[i]] = record{true, i}
	}

	for i = range nums {
		if m[target-nums[i]].found {
			break
		}
	}

	return []int{i, m[target-nums[i]].idx}
}

func main() {
	fmt.Printf("twoSum([]int{2, 7, 11, 15}, 9): %v\n", twoSum([]int{2, 7, 11, 15}, 9))
}
