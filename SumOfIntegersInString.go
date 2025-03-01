//Your task in this kata is to implement a function that calculates the sum of the integers inside a string. For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.

package kata

import (
	"regexp"
	"strconv"
)

func SumOfIntegersInString(strng string) (res int) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(strng, -1)

	for _, match := range matches {
		intValue, _ := strconv.Atoi(match)
		res += intValue
	}

	return res
}
