/*
Simple, given a string of words, return the length of the shortest word(s).

String will never be empty and you do not need to account for different data types.
*/

package kata

import "strings"

func FindShort(s string) (res int) {
	a := strings.Split(s, " ")
	res = 1000
	for _, word := range a {
		if len(word) < res {
			res = len(word)
		}
	}
	return
}

/*
Best solution
package kata

import "strings"

func FindShort(s string) int {
  shortest := len(s)
  for _, word := range strings.Split(s, " ") {
    if len(word) < shortest {
      shortest = len(word)
    }
  }
  return shortest
}
*/
