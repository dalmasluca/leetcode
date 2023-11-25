package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	b := make([]bool, len(candies))
	var max int = candies[0]
	for i := 1; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	for i := range candies {
		if candies[i]+extraCandies >= max {
			b[i] = true
		} else {
			b[i] = false
		}
	}
	return b
}
