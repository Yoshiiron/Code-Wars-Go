//You might know some pretty large perfect squares. But what about the NEXT one?

package kata

import (
	"math"
)

func FindNextSquare(sq int64) int64 {
	root := int(math.Sqrt(float64(sq)))
	if int64(root*root) != sq {
		return -1
	}
	nextRoot := root + 1
	return int64(nextRoot * nextRoot)
}

/*
package kata

import (
  "math"
)

func FindNextSquare(sq int64) int64 {
  res := math.Pow(math.Sqrt(float64(sq)) + 1, 2)

  if res == math.Trunc(res) {
    return int64(res)
  }
  return -1
}
*/
