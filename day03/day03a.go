package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Read file

func main() {
	powerfile, err := os.Open("./powerlevels.txt")
	chkerr(err)
	defer powerfile.Close()

	powerlevels := bufio.NewScanner(powerfile)

	var powerlines []string
	for powerlevels.Scan() {
		powerlines = append(powerlines, powerlevels.Text())
	}

	var gammacount [12]int
	var epsiloncount [12]int

	for _, level := range powerlines {

		powerbin, err := strconv.ParseUint(level, 2, 64)
		chkerr(err)

		for i := 0; i < 12; i++ {
			var bitlevel uint64 = 1 << i

			if powerbin&bitlevel == bitlevel {
				gammacount[i]++
			} else {
				epsiloncount[i]++
			}
		}
	}

	var gamma int
	var epsilon int

	for i := 0; i < 12; i++ {
		if gammacount[i] > epsiloncount[i] {
			gamma += int(math.Pow(2, float64(i)))
		} else {
			epsilon += int(math.Pow(2, float64(i)))
		}
	}

	fmt.Println(gamma * epsilon)
}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
