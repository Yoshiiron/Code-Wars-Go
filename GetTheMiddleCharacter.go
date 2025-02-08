/*
You are going to be given a non-empty string. Your job is to return the middle character(s) of the string.
If the string's length is odd, return the middle character.
If the string's length is even, return the middle 2 characters.
*/

package kata

import (
	"math"
	"strings"
)

func GetMiddle(s string) (res string) {
	splited := strings.Split(s, "")
	b := int(math.Ceil(float64(len(splited)) / 2))
	res += splited[b-1]
	if (len(splited) % 2) == 0 {
		res += splited[b]
		return
	} else {
		return
	}
}

/*
Best solution

package kata

func GetMiddle(s string) string {
    n := len(s)
    if n==0 {return s}
    return s[(n - 1) / 2 : n / 2 + 1]
}
*/
