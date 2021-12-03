package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Read file

func main() {
	depthsfile, err := os.Open("./depths.txt")
	chkerr(err)
	defer depthsfile.Close()

	depths := bufio.NewScanner(depthsfile)
	depths.Split(bufio.ScanLines)

	var depthslines []string

	for depths.Scan() {
		depthslines = append(depthslines, depths.Text())
	}

	var increasecount int
	var previousdepth int
	var previousdepthtrio int

	depthfirst, err := strconv.Atoi(depthslines[0])
	chkerr(err)
	depthsecond, err := strconv.Atoi(depthslines[1])
	chkerr(err)
	depththird, err := strconv.Atoi(depthslines[2])
	chkerr(err)

	increasecount = 0
	previousdepth = depthfirst
	previousdepthtrio = depthfirst + depthsecond + depththird

	fmt.Println(previousdepthtrio)

	for i, depth := range depthslines {

		if i == 0 || i == 1 || i == 2 {
			continue
		}

		depthint, err := strconv.Atoi(depth)
		chkerr(err)
		fmt.Println(depthint)

		depthfirst = depthsecond
		depthsecond = depththird
		depththird = depthint

		newdepthtrio := previousdepthtrio + depthint - previousdepth
		fmt.Println(newdepthtrio)

		if newdepthtrio > previousdepthtrio {
			increasecount++
		}

		fmt.Println(increasecount)
		previousdepthtrio = newdepthtrio
		previousdepth = depthfirst

		fmt.Println(previousdepthtrio)
		fmt.Println(" ")
	}

	fmt.Println(increasecount)

}

// for loop through lines checking if greater than previous

// what is starting value
// how best to store current value between iterations

// print answer

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
