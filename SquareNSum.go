/*
Complete the square sum function so that it squares each number passed into it and then sums the results together.
*/

package kata

func SquareSum(numbers []int) (res int) {
	for _, num := range numbers {
		res += num * num
	}
	return res
}
