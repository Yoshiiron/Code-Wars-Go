/*
Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer.

Square all numbers k (0 <= k <= n) between 0 and n.

Count the numbers of digits d used in the writing of all the k**2.

Implement the function taking n and d as parameters and returning this count.
*/

package kata

import (
	"strconv"
)

func NbDig(n int, d int) (count int) {
	targetDigit := strconv.Itoa(d)
	for k := 0; k <= n; k++ {
		squaredNum := k * k
		squaredStr := strconv.Itoa(squaredNum)
		for _, digit := range squaredStr {
			if string(digit) == targetDigit {
				count++
			}
		}

	}
	return count
}

/*
package kata
import (
  "strings"
  "strconv"
)

func NbDig(n int, d int) (out int) {
  for i := 0; i <= n; i++ {
    out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
  }

  return
}
*/
