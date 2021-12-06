package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read file

func main() {
	movementfile, err := os.Open("./movements.txt")
	chkerr(err)
	defer movementfile.Close()

	movements := bufio.NewScanner(movementfile)
	movements.Split(bufio.ScanLines)

	var movementlines []string

	for movements.Scan() {
		movementlines = append(movementlines, movements.Text())
	}

	var forward int
	var depth int

	forward = 0
	depth = 0

	for _, move := range movementlines {

		// Split into type and amount

		words := strings.Fields(move)

		movetype := words[0]
		moveamountint, err := strconv.Atoi(words[1])
		chkerr(err)

		switch movetype {
		case "forward":
			forward += moveamountint
		case "down":
			depth += moveamountint
		case "up":
			depth -= moveamountint
		}
	}

	fmt.Println(forward * depth)

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
