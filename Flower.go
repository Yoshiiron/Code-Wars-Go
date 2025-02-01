/*
Who remembers back to their time in the schoolyard, when girls would take a flower and tear its petals, saying each of the following phrases each time a petal was torn:

"I love you"
"a little"
"a lot"
"passionately"
"madly"
"not at all"
If there are more than 6 petals, you start over with "I love you" for 7 petals, "a little" for 8 petals and so on.

*/

package kata

func HowMuchILoveYou(i int) string {
	for a := i; a > 6; a -= 6 {
		i = i - 6
	}
	switch i {
	case 1:
		return "I love you"
	case 2:
		return "a little"
	case 3:
		return "a lot"
	case 4:
		return "passionately"
	case 5:
		return "madly"
	case 6:
		return "not at all"
	}
	return ""
}

/*
package kata

func HowMuchILoveYou(i int) string {
    return []string{"I love you","a little","a lot","passionately","madly","not at all"}[(i-1) % 6]
}
*/
