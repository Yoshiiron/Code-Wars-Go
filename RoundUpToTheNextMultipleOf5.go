//Given an integer as input, can you round it to the next (meaning, "greater than or equal") multiple of 5?

package kata

func RoundToNext5(n int) int {
	for {
		if n%5 != 0 {
			n += 1
		} else {
			break
		}
	}
	return n
}

/*
Best solution
package kata


func RoundToNext5(n int) int {
  for n % 5 != 0 {
    n++
  }
  return n
}
*/
