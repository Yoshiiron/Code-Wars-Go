//In a small town the population is p0 = 1000 at the beginning of a year. The population regularly increases by 2 percent per year and moreover 50 new inhabitants per year come to live in the town. How many years does the town need to see its population greater than or equal to p = 1200 inhabitants?

package kata

func NbYear(p0 int, percent float64, aug int, p int) (res int) {
	for {
		if p0 < p {
			p0 = int(float64(p0) + float64(p0)*(percent/100) + float64(aug))
		} else {
			break
		}
		res += 1
	}
	return
}
