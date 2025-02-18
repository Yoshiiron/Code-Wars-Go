/*
You are given two sorted arrays that contain only integers. These arrays may be sorted in either ascending or descending order. Your task is to merge them into a single array, ensuring that:

The resulting array is sorted in ascending order.

Any duplicate values are removed, so each integer appears only once.

If both input arrays are empty, return an empty array.

No input validation is needed, as both arrays are guaranteed to contain zero or more integers.
*/

package kata

import "sort"

func MergeArrays(arr1, arr2 []int) (result []int) {
	merged := append(arr1, arr2...)
	
	sort.Ints(merged)

	for ind, digit := range merged {
		if 
	}

	return result
}
