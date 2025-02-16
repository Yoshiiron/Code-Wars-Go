// Given a string s, your task is to return another string such that even-indexed and odd-indexed characters of s are grouped and the groups are space-separated. Even-indexed group comes as first, followed by a space, and then by the odd-indexed part.

package kata

func SortMyString(s string) string {
	var odd, even string
	for ind, str := range s {
		if ind%2 != 0 {
			odd += string(str)
		} else {
			even += string(str)
		}
	}

	result := even + " " + odd

	return result

}
