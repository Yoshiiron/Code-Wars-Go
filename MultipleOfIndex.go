//Return a new array consisting of elements which are multiple of their own index in input array (length > 1).

package kata

func multipleOfIndex(ints []int) (res []int) {
	for index, inte := range ints {
		if index != 0 && inte%index == 0 {
			res = append(res, inte)
		}
	}
	return
}
