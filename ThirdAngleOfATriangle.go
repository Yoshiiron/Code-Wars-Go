/*
You are given two interior angles (in degrees) of a triangle.

Write a function to return the 3rd.
*/

package kata

func OtherAngle(a int, b int) int {
	return 180 - (a + b)
}
