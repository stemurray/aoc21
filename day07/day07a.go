package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Open file and read as words without whitespace
	crabfile, err := os.Open("./crabs.txt")
	chkerr(err)
	defer crabfile.Close()

	crabscan := bufio.NewScanner(crabfile)
	crabscan.Split(bufio.ScanWords)

	// Read first line (which is comma separated) and add to a slice of strings
	crabscan.Scan()
	crabstr := crabscan.Text()
	crabslc := strings.Split(crabstr, ",")
	var crabs []int

	for _, c := range crabslc {
		crabint, err := strconv.Atoi(c)
		chkerr(err)
		crabs = append(crabs, crabint)
	}

	fmt.Println(crabs)

	var lowestinitialposition int
	var highestinitialposition int

	for _, p := range crabs {
		if p < lowestinitialposition {
			lowestinitialposition = p
		}
		if p > highestinitialposition {
			highestinitialposition = p
		}
	}

	var fuel int

	for _, p := range crabs {
		fuel = fuel + Abs(lowestinitialposition-p)
	}

	for t := lowestinitialposition; t < highestinitialposition; t++ {

		var f int

		for _, p := range crabs {
			f = f + Abs(t-p)
		}

		if f < fuel {
			fuel = f
		}
	}

	fmt.Println(fuel)

}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
