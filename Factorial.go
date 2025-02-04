//Your task is to write function factorial.

package kata

func Factorial(n int) (res int) {
	res += 1
	for i := n; i > 0; i-- {
		res *= i
	}
	return
}
