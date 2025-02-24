package kata

import "strconv"

func StringToNumber(str string) (res int) {
	res, _ = strconv.Atoi(str)
	return
}
