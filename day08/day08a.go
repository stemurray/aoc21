package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Open file and read as words without whitespace
	signalfile, err := os.Open("./signals.txt")
	chkerr(err)
	defer signalfile.Close()

	signalscan := bufio.NewScanner(signalfile)
	signalscan.Split(bufio.ScanLines)

	// Read first line (which is comma separated) and add to a slice of strings
	var count int
	for signalscan.Scan() {
		signalstr := signalscan.Text()
		signalslc := strings.Split(signalstr, " ")

		for i := 11; i < 15; i++ {
			if len(signalslc[i]) == 2 || len(signalslc[i]) == 3 || len(signalslc[i]) == 4 || len(signalslc[i]) == 7 {
				count++
			}
		}

	}

	fmt.Println(count)

}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
