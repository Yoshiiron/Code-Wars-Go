//Write a function named setAlarm/set_alarm/set-alarm/setalarm (depending on language) which receives two parameters. The first parameter, employed, is true whenever you are employed and the second parameter, vacation is true whenever you are on vacation.

package kata

func SetAlarm(employed, vacation bool) bool {
	if employed && !vacation {
		return true
	}
	return false
}
