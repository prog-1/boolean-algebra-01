package main

import "fmt"

func implies(a, b bool) bool {
	if !a {
		return true
	}
	return false
}

func proposition(a, b, c, d bool) bool {
	return implies(a, xor(b, c)) && implies(!b, a && d) && implies(d, xor(b, c))
}

func xor(b, c bool) bool {
	return b != c
}

func main() {
	fmt.Println(`A: If Jones did not meet Smith last night 
B: then either Smith was a murderer
C: or Jones is telling a lie If Smith was not a murderer 
D: then Jones did not meet Smith last night, and the murder happened after midnight 

If the murder happened after midnight, then either Smith was a murderer, or Jones is telling a lie, but not both.`)
	fmt.Println(" ")

	bools := [2]bool{false, true}
	fmt.Println("    A      B      C      D      (A => (B (+) C)) && (!B => (A & D)) && (D => (B (+) C))")
	for _, a := range bools {
		for _, b := range bools {
			for _, c := range bools {
				for _, d := range bools {
					fmt.Printf("%7v%7v%7v%7v%33v\n", a, b, c, d, proposition(a, b, c, d))
				}
			}
		}
	}
}
