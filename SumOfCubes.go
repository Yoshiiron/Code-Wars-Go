// Write a function that takes a positive integer n, sums all the cubed values from 1 to n (inclusive), and returns that sum.

package kata

func SumCubes(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i * i * i
	}
	return
}
