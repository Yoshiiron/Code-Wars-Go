//Implement a function that computes the difference between two lists. The function should remove all occurrences of elements from the first list (a) that are present in the second list (b). The order of elements in the first list should be preserved in the result.

package kata

func ArrayDiff(a, b []int) (res []int) {
	m := make(map[int]struct{})

	for _, num := range b {
		m[num] = struct{}{}
	}

	for _, val := range a {
		if _, found := m[val]; !found {
			res = append(res, val)
		}
	}

	return res

}
