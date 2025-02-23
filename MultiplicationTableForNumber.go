//Your goal is to return multiplication table for number that is always an integer from 1 to 10.

package kata

import (
	"fmt"
	"strings"
)

func MultiTable(number int) (res string) {
	sl := []string{}
	for i := 1; i <= 10; i++ {
		sl = append(sl, fmt.Sprintf("%v * %v = %v", i, number, i*number))
	}
	return strings.Join(sl, "\n")
}
