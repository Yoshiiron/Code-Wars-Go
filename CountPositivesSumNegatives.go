//Your task is to sort a given string. Each word in the string will contain a single number. This number is the position the word should have in the result.

package kata

func CountPositivesSumNegatives(numbers []int) (res []int) {
	pos := 0
	neg := 0

	for _, dig := range numbers {
		if dig > 0 {
			pos++
		} else if dig < 0 {
			neg += dig
		}
	}
	res = append(res, pos, neg)

	return res // your code here
}

/*
package kata

func CountPositivesSumNegatives(numbers []int) []int {
  res := []int{0,0}
  for _, v := range numbers {
    if v > 0 {
      res[0] += 1
    }else {
      res[1] += v
    }
  }
  return res // your code here
}
*/
