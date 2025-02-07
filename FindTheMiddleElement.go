/*
As a part of this Kata, you need to create a function that when provided with a triplet, returns the index of the numerical element that lies between the other two elements.

The input to the function will be an array of three distinct numbers (Haskell: a tuple).
*/

package kata

import (
	"sort"
)

func Gimme(array [3]int) (res int) {
	sortSlice := array[:]
	slice := make([]int, len(sortSlice))
	copy(slice, sortSlice)
	sort.Ints(sortSlice)

	findThis := sortSlice[1]
	for ind, dig := range slice {
		if dig == findThis {
			res = ind
		}
	}

	return res
}
