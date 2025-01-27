//In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". Your function receives one side of the DNA (string, except for Haskell); you need to return the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).

package kata

func DNAStrand(dna string) (res string) {
	for _, run := range dna {
		switch string(run) {
		case "A":
			res += "T"
		case "T":
			res += "A"
		case "C":
			res += "G"
		case "G":
			res += "C"
		}
	}
	return res
}

/*
Best solutions

package kata

import "strings"

var dnaReplacer *strings.Replacer = strings.NewReplacer(
  "A", "T",
  "T", "A",
  "C", "G",
  "G", "C",
)

func DNAStrand(dna string) string {
  return dnaReplacer.Replace(dna)
}


package kata

import (
    "strings"
)

func DNAStrand(dna string) string {
  // your code here
  replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
	return(replacer.Replace(dna))
}

*/
