//Complete the function which converts a binary number (given as a string) to a decimal number.

package kata

import "strconv"

func BinToDec(bin string) int {
	dig, _ := strconv.ParseInt(bin, 2, 64)
	return int(dig)
}
