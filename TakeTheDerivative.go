/*
This function takes two numbers as parameters, the first number being the coefficient, and the second number being the exponent.

Your function should multiply the two numbers, and then subtract 1 from the exponent. Then, it has to return an expression (like 28x^7). "^1" should not be truncated when exponent = 2.
*/

package kata

import "fmt"

func Derive(coefficient, exponent int) string {
	return fmt.Sprintf("%dx^%d", coefficient*exponent, exponent-1)
}
