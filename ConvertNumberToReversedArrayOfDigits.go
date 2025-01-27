//Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

package kata

import (
	"strconv"
)

func Digitize(n int) (res []int) {
	numStr := strconv.Itoa(n)
	for i := len(numStr) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(numStr[i]))
		res = append(res, digit)
	}
	return res
}
