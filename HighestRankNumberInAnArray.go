/*
Complete the method which returns the number which is most frequent in the given input array. If there is a tie for most frequent number, return the largest number among them.

Note: no empty arrays will be given.
*/

package kata

func HighestRank(nums []int) (res int) {

	dig := make(map[int]int)

	for _, num := range nums {
		_, ok := dig[num]
		if !ok {
			dig[num] = 1
		} else {
			dig[num] += 1
		}
	}

	highEstests := []int{}
	high := 0

	for key := range dig {
		if high < dig[key] {
			high = dig[key]
		}
	}

	for key := range dig {
		if dig[key] == high {
			highEstests = append(highEstests, key)
		}
	}

	for _, num := range highEstests {
		if num > res {
			res = num
		}
	}

	return
}

/*
Best solution


package kata

func HighestRank(nums []int) int {
  mii, maxK, maxV := map[int]int{}, 0, 0
  for _, v := range nums {
    mii[v]++
    if mii[v] > maxV || (mii[v] == maxV && v > maxK) {
      maxK = v
      maxV = mii[v]
    }
  }
  return maxK
}

*/
