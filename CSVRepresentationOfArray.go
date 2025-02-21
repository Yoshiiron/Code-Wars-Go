//Create a function that returns the CSV representation of a two-dimensional numeric array.

package kata

import (
	"strconv"
)

func ToCsvText(array [][]int) (res string) {
	for j, arr := range array {
		for i, num := range arr {
			if len(arr) == (i + 1) {
				res += strconv.Itoa(num)
			} else {
				res += strconv.Itoa(num) + ","
			}
		}
		if len(array) != (j + 1) {
			res += "\n"
		}
	}
	return res
}

/*
Best solution

package kata

import (
  "fmt"
  "strings"
)

func ToCsvText(array [][]int) string {

  lines := []string{}

  for _, line := range array {
    new_line := strings.Trim(fmt.Sprint(line), "[]")
    new_line = strings.ReplaceAll(new_line, " ", ",")
    lines = append(lines, new_line)
  }

  return strings.Join(lines, "\n")
}

*/
