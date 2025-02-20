// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

package kata

import (
	"strings"
)

func duplicate_count(s1 string) (i int) {
	ma := make(map[rune]int)

	s1 = strings.ToLower(s1)

	for _, run := range s1 {
		if _, exists := ma[run]; !exists {
			ma[run] += 1
		} else {
			ma[run] += 1
		}
	}

	for key := range ma {
		if ma[key] > 1 {
			i++
		}
	}

	return i //Instead of returning '1', you have to return integer 'i' as answer of solution.
}

/*
Best solution

package kata
import "strings"
func duplicate_count(s string) (c int) {
    h := map[rune]int{}
    for _, r := range strings.ToLower(s) {
      if h[r]++; h[r] == 2 { c++ }
    }
    return
}

*/
