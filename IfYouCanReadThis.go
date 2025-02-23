//You'll have to translate a string to Pilot's alphabet (NATO phonetic alphabet).

package kata

import (
	"strings"
	"unicode"
)

func ToNato(words string) (res string) {
	// the NATO map[string]string is preloaded for you
	// NATO["A"] == "Alfa", etc
	var NATO = map[string]string{
		"A": "Alfa",
		"B": "Bravo",
		"C": "Charlie",
		"D": "Delta",
		"E": "Echo",
		"F": "Foxtrot",
		"G": "Golf",
		"H": "Hotel",
		"I": "India",
		"J": "Juliett",
		"K": "Kilo",
		"L": "Lima",
		"M": "Mike",
		"N": "November",
		"O": "Oscar",
		"P": "Papa",
		"Q": "Quebec",
		"R": "Romeo",
		"S": "Sierra",
		"T": "Tango",
		"U": "Uniform",
		"V": "Victor",
		"W": "Whiskey",
		"X": "Xray",
		"Y": "Yankee",
		"Z": "Zulu",
	}
	words = strings.ReplaceAll(words, " ", "")
	for _, word := range words {
		if unicode.IsPunct(word) {
			res += string(word)
		} else {
			res += NATO[strings.ToUpper(string(word))] + " "
		}
	}
	return strings.TrimSpace(res)
}
