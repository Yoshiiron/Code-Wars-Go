/*
Make a function that will return a greeting statement that uses an input; your program should return, "Hello, <name> how are you doing today?".

[Make sure you type the exact thing I wrote or the program may not execute properly]
*/

package kata

import "fmt"

func Greet(name string) string {
	// fill in solution here
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}
