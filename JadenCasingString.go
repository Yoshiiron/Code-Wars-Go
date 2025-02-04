//Your task is to convert strings to how they would be written by Jaden Smith. The strings are actual quotes from Jaden Smith, but they are not capitalized in the same way he originally typed them.

package kata

import "strings"

func ToJadenCase(str string) string {
	return strings.Title(str)
}
