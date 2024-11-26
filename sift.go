package main

import (
	"fmt"
)

const sm = 32768
const lg = sm * sm

var s [lg]int8

func prune(p int) {
	p2 := p + p
	for i := p * p; i < lg; i += p2 {
		s[i] = 0
	}
}

func main() {
	// fmt.Printf("%v threads available.\n", runtime.GOMAXPROCS(-1))
	primes := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73}

	s[2] = 1
	for i := 3; i < lg; i += 2 {
		s[i] = 1
	}

	for i := 0; i < len(primes); i++ {
		go prune(primes[i])
	}

	next := 74 // started at 3
	for next < sm {
		for ; s[next] == 0; next++ {
		}
		prune(next)
		next++
	}

	count := 0
	for i := 0; i < lg; i++ {
		if s[i] == 1 {
			count++
		}
	}
	fmt.Println("Found ", count, "Primes between 0 and ", lg)
	for i := 0; i < 102; i++ {
		if s[i] == 1 {
			fmt.Print(i, ", ")
		}
	}
	fmt.Println("")
	for i := 120; i > 1; i-- {
		if s[lg-i] == 1 {
			fmt.Print(lg-i, " ")
		}
	}
	fmt.Println("")
}
