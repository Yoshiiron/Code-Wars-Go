/*
Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.
*/

package kata

func Invert(arr []int) (res []int) {
	for _, b := range arr {
		res = append(res, b*-1)
	}
	return
}
