/*
Kata Task
I have a cat and a dog.

I got them at the same time as kitten/puppy. That was humanYears years ago.

Return their respective ages now as [humanYears,catYears,dogYears]

NOTES:

humanYears >= 1
humanYears are whole numbers only
Cat Years
15 cat years for first year
+9 cat years for second year
+4 cat years for each year after that
Dog Years
15 dog years for first year
+9 dog years for second year
+5 dog years for each year after that
*/

package kata

func CalculateYears(years int) (result [3]int) {
	result[0] = years
	cat := 0
	dog := 0
	for i := years; i > 0; i-- {
		switch i {
		case 1:
			cat += 15
			dog += 15
		case 2:
			cat += 9
			dog += 9
		default:
			cat += 4
			dog += 5
		}
	}
	result[1], result[2] = cat, dog
	return
}

/*
Best solution

package kata

func CalculateYears(years int) (result [3]int) {
  switch years {
    case 1: result = [3]int{1, 15, 15}
    case 2: result = [3]int{2, 24, 24}
    default: result = [3]int{years, 24 + 4 * (years - 2), 24 + 5 * (years - 2)}
  }
  return
}
*/
