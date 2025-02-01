//Build a function that returns an array of integers from n to 1 where n>0.

package kata

func ReverseSeq(n int) (res []int) {
	for i := n; i > 0; i-- {
		res = append(res, i)
	}
	return
}
