//Create a combat function that takes the player's current health and the amount of damage received, and returns the player's new health. Health can't be less than 0.

package kata

func combat(health, damage float64) float64 {
	if health-damage < 0 {
		return 0
	}
	return health - damage
}

/*

Best solution

package kata
import "math"
func combat(health, damage float64) float64 {
  return math.Max(health-damage, 0.0)
}

*/
