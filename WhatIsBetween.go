/*
Complete the function that takes two integers (a, b, where a < b) and return an array of all integers between the input parameters, including them.
*/

package kata

func Between(a, b int) (res []int) {
	for i := a; i <= b; i++ {
		res = append(res, i)
	}
	return res
}
