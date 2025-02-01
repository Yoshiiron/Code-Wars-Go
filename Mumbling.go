/*
This time no story, no theory. The examples below show you how to write function accum:

Examples:
accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
*/

package kata

import (
	"strings"
)

func Accum(s string) (res string) {
	for index, letter := range s {
		for i := index + 1; i > 0; i-- {
			if i == index+1 {
				res += strings.ToUpper(string(letter))
			} else {
				res += strings.ToLower(string(letter))
			}
		}
		if index != len(s)-1 {
			res += "-"
		}
	}
	return
}

/*Best solution
package kata

import "strings"

func Accum(s string) string {
    parts := make([]string, len(s))

    for i := 0; i < len(s); i++ {
      parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
    }

    return strings.Join(parts, "-")
}
*/
