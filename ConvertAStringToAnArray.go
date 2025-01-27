//Write a function to split a string and convert it into an array of words.

package kata

import (
	"strings"
)

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}

/*
Best solution
package kata

import "strings"

func StringToArray(str string) []string {
      return strings.Fields(str)
}
*/
