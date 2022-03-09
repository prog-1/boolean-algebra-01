package main

import "fmt"

func xor(a, b bool) bool {
	return a != b
}

func implies(a, b bool) bool {
	return !a || b
}

func proposition(a, b, c, d bool) bool {
	return (implies(a, xor(b, c))) && (implies(!b, a && d)) && (implies(d, xor(b, c)))
}

func main() {
	bools := []bool{false, true}
	fmt.Println("    A      B      C      D		(A => (B (+) C)) && (!B => (A & D)) && (D => (B (+) C))")
	for _, a := range bools {
		for _, b := range bools {
			for _, c := range bools {
				for _, d := range bools {
					fmt.Printf("%7v%7v%7v%7v%20v\n", a, b, c, d, proposition(a, b, c, d))
				}
			}
		}
	}
}
