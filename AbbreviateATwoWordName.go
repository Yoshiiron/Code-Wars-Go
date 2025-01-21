//Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
//The output should be two capital letters with a dot separating them.

package kata

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	strin := strings.Split(name, " ")
	nameA := strings.Split(strin[0], "")
	lastnameA := strings.Split(strin[1], "")
	return fmt.Sprintf("%s.%s", strings.ToTitle(nameA[0]), strings.ToTitle(lastnameA[0]))
}

/* Best solution
func AbbrevName(name string) string {
	strin := strings.Split(name, " ")
	return fmt.Sprintf("%s.%s", strings.ToTitle(string(strin[0][0])), strings.ToTitle(string(strin[0][0])))
}
*/
