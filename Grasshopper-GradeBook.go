//Complete the function so that it finds the average of the three scores passed to it and returns the letter value associated with that grade.

package kata

func GetGrade(a, b, c int) rune {
	aver := (a + b + c) / 3

	switch {
	case 90 <= aver || aver <= 100:
		return 'A'
	case 80 <= aver || aver < 90:
		return 'B'
	case 70 <= aver || aver < 80:
		return 'C'
	case 60 <= aver || aver < 70:
		return 'D'
	case 0 <= aver || aver < 60:
		return 'F'
	}

	return ' '

}
