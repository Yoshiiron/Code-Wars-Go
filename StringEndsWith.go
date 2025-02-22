// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

package kata

import (
	"strings"
)

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
