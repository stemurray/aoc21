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
	fishfile, err := os.Open("./fish.txt")
	chkerr(err)
	defer fishfile.Close()

	fishscan := bufio.NewScanner(fishfile)
	fishscan.Split(bufio.ScanWords)

	// Read first line (which is comma separated) and add to a slice of strings
	fishscan.Scan()
	fishstr := fishscan.Text()
	fishesstr := strings.Split(fishstr, ",")
	var fishes [9]int

	for _, f := range fishesstr {
		fishint, err := strconv.Atoi(f)
		chkerr(err)
		fishes[fishint]++
	}

	for i := 0; i < 256; i++ {
		fishzero := fishes[0]
		fmt.Println(fishes)
		for days, fish := range fishes {
			switch days {
			case 8, 6, 5, 4, 3, 2, 1:
				fishes[days-1] = fish
			case 7:
				fishes[6] = fishzero + fish
			case 0:
				fishes[8] = fish
			}
		}
	}

	var total int
	for _, f := range fishes {
		total = total + f
	}

	fmt.Println(total)
}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
