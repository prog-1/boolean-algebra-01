package main

import (
	"fmt"
)

func state(a, b bool) bool {
	a = !a
	return a || b
}
func proposition(a, b, c, d bool) bool {
	return (state(a, b != c)) && (state(d, b != c)) && (state(!b, a && d))
}
func main() {
	bools := [2]bool{false, true}
	fmt.Println("A       B       C       D       (A => (B^C)) & (!B=>(A&D)) & (D=>(B^C)) ")
	for _, a := range bools {
		for _, b := range bools {
			for _, c := range bools {
				for _, d := range bools {
					fmt.Printf("%3v%8v%8v%8v%10v\n", a, b, c, d, proposition(a, b, c, d))
				}

			}
		}
	}
}
