//Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

//Return your answer as a number.

package kata

import "strconv"

func SumMix(arr []any) (res int) {
	for _, sumi := range arr {
		switch sumi.(type) {
		case int:
			res += sumi.(int)
		case string:
			conv, _ := strconv.Atoi(sumi.(string))
			res += conv
		}
	}
	return
}
