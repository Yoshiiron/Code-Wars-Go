/*
Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.
*/

package kata

import (
	"strings"
)

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseWords(str string) string {

	wordsSl := strings.Split(str, " ")
	for i, words := range wordsSl {
		wordsSl[i] = reverseString(words)
	}

	return strings.Join(wordsSl, " ")
}

/*
Best solution

package kata

import "strings"

func Reverse(str string) (result string) {
	for _, v := range str{
		result = string(v) + result
	}
	return result
}

func ReverseWords(str string) (result string) {
	words := strings.Split(str, " ")
	for i, v := range words{
		words[i] = Reverse(string(v))
	}

	return strings.Join(words, " ") // reverse those words
}
*/
