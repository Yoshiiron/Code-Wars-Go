//You get an array of numbers, return the sum of all of the positives ones.

package kata

func PositiveSum(numbers []int) (res int) {
	for _, in := range numbers {
		if in > 0 {
			res += in
		}
	}
	return
}
