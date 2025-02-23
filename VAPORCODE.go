//Write a function that converts any sentence into a V A P O R W A V E sentence. a V A P O R W A V E sentence converts all the letters into uppercase, and adds 2 spaces between each letter (or special character) to create this V A P O R W A V E effect.

package kata

import "strings"

func Vaporcode(s string) (res string) {
	str := s
	str = strings.ReplaceAll(str, " ", "")
	for _, run := range str {
		res += strings.ToUpper(string(run)) + "  "
	}
	res = strings.TrimRight(res, "  ")
	return
}

/*
Best solution

package kata

import "strings"

func Vaporcode(s string) string {
  return strings.Join(strings.Split((strings.ReplaceAll(strings.ToUpper(s), " ", "")), ""), "  ")
}
*/
