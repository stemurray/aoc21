package main

import (
	"bufio"
	"fmt"
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

	// GAMMA
	gammaslice := append([]string(nil), powerlines...)

	i := 11
	for len(gammaslice) > 1 {

		var tempgamma []string
		var gammacount int
		var gammanum int

		for _, level := range gammaslice {

			powerbin, err := strconv.ParseUint(level, 2, 64)
			chkerr(err)

			var bitlevel uint64 = 1 << i

			if powerbin&bitlevel == bitlevel {
				gammacount++
			}
		}

		fmt.Println(gammacount)
		fmt.Println(len(gammaslice))

		if gammacount < (len(gammaslice) - gammacount) {
			gammanum = 0
		} else {
			gammanum = 1
		}
		fmt.Println(gammanum)
		fmt.Println(" ")

		for _, level := range gammaslice {

			powerbin, err := strconv.ParseUint(level, 2, 64)
			chkerr(err)

			var bitlevel uint64 = 1 << i

			if gammanum == 1 {
				if powerbin&bitlevel == bitlevel {
					tempgamma = append(tempgamma, level)
				}
			} else {
				if powerbin&bitlevel != bitlevel {
					tempgamma = append(tempgamma, level)
				}
			}
		}

		gammaslice = append([]string(nil), tempgamma...)
		i--
		fmt.Println(gammaslice)
		fmt.Println(i)
	}

	// EPSILON
	epsslice := append([]string(nil), powerlines...)

	j := 11
	for len(epsslice) > 1 {

		var tempeps []string
		var epscount int
		var epsnum int

		for _, level := range epsslice {

			powerbin, err := strconv.ParseUint(level, 2, 64)
			chkerr(err)

			var bitlevel uint64 = 1 << j

			if powerbin&bitlevel == bitlevel {
				epscount++
			}
		}
		fmt.Println(epscount)
		fmt.Println(len(epsslice))

		if epscount < (len(epsslice) - epscount) {
			epsnum = 1
		} else {
			epsnum = 0
		}
		fmt.Println(epsnum)
		fmt.Println(" ")

		for _, level := range epsslice {

			powerbin, err := strconv.ParseUint(level, 2, 64)
			chkerr(err)

			var bitlevel uint64 = 1 << j

			if epsnum == 1 {
				if powerbin&bitlevel == bitlevel {
					tempeps = append(tempeps, level)
				}
			} else {
				if powerbin&bitlevel != bitlevel {
					tempeps = append(tempeps, level)
				}
			}
		}

		epsslice = append([]string(nil), tempeps...)
		j--
		fmt.Println(epsslice)
		fmt.Println(j)
	}

	fmt.Println(gammaslice[0])
	fmt.Println(epsslice[0])
	fmt.Println(" ")

	var gamma int
	var epsilon int

	gbin, err := strconv.ParseUint(gammaslice[0], 2, 64)
	gamma = int(gbin)

	ebin, err := strconv.ParseUint(epsslice[0], 2, 64)
	epsilon = int(ebin)

	fmt.Println(gamma)
	fmt.Println(epsilon)
	fmt.Println(gamma * epsilon)

}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
