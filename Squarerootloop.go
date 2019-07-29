package main

import (
	"fmt"
)

var afl = [4]float64{1, 2, 3, 4}

//8 characters long using starting at -1 and going to -8

func squareroot(x float64) float64 {
	z := 1.0
	con := false
	var eplace int
	for index := 0; con != true; index++ {
		z -= (z*z - x) / (2 * z)
		afl[eplace] = z
		count := 0
		for i := range afl {
			if afl[i] == z {
				count++
			} else {
				count = 0
				break

			}
			if count == 4 {
				con = true
			}
		}
		fmt.Printf("round: %v\n", index+1)
		fmt.Println(z)

		if eplace < 3 {
			eplace++
		} else {
			eplace = 0
		}

	}
	return z
}

func main() {
	squareroot(333333874390895374209857304982502)
}
