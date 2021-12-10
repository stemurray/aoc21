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
	var fishes []int

	for _, f := range fishesstr {
		fishint, err := strconv.Atoi(f)
		chkerr(err)
		fishes = append(fishes, fishint)
	}

	fmt.Println(fishes)

	for i := 0; i < 80; i++ {
		for fish, days := range fishes {
			switch days {
			case 8, 7, 6, 5, 4, 3, 2, 1:
				fishes[fish]--
			case 0:
				fishes[fish] = 6
				fishes = append(fishes, 8)
			}
		}
	}

	fmt.Println("")
	fmt.Println(fishes)
	fmt.Println(len(fishes))
}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
