//Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase.

package kata

import (
	"unicode"
)

func ToAlternatingCase(str string) (res string) {

	for _, run := range str {
		if unicode.IsLower(run) {
			res += string(unicode.ToUpper(run))
			continue
		} else if unicode.IsUpper(run) {
			res += string(unicode.ToLower(run))
			continue
		}
		res += string(run)
	}

	return res
}
