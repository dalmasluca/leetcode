package main

func romanToInt(s string) int {

	romanToDecimalMap := make(map[byte]int)

	romanToDecimalMap['I'] = 1
	romanToDecimalMap['V'] = 5
	romanToDecimalMap['X'] = 10
	romanToDecimalMap['L'] = 50
	romanToDecimalMap['C'] = 100
	romanToDecimalMap['D'] = 500
	romanToDecimalMap['M'] = 1000

	result := 0

	for i := 0; i < len(s)-1; i++ {
		val := romanToDecimalMap[s[i]]
		if val < romanToDecimalMap[s[i+1]] {
			result -= val
		} else {
			result += val
		}
	}
	result += romanToDecimalMap[s[len(s)-1]]
	return result
}
