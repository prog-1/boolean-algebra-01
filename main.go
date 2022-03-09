package main

import "fmt"

func implies(a, b bool) bool {
	return !a || b
}

func proposition(a, b, c, d bool) bool {
	//fmt.Println(implies(a, b || c), implies(!b, a && d), implies(d, b != c))
	return (implies(a, b || c) && implies(!b, a && d) && implies(d, b != c))
}

func main() {
	/*
		A - Jones did not meet Smith last night
		B - Smith was a murderer
		C - Jones is telling a lie
		D - Murder happened after midnight
	*/
	bools := [2]bool{false, true}
	fmt.Println("      A      B      C      D      (a=>b || c) && (!b => a && b) && (d => b xor c)")
	for _, a := range bools {
		for _, b := range bools {
			for _, c := range bools {
				for _, d := range bools {
					fmt.Printf("%7v%7v%7v%7v%32v\n", a, b, c, d, proposition(a, b, c, d))
				}
			}
		}
	}
}
