// Write a function that removes the spaces from the string, then return the resultant string.

package kata

import "strings"

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}
