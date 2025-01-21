// Consider an array/list of sheep where some sheep may be missing from their place. We need a function that counts the number of sheep present in the array (true means present).
package kata

func CountSheeps(numbers []bool) int {
	var result int
	for _, state := range numbers {
		if state {
			result++
		}
	}
	return result
}
