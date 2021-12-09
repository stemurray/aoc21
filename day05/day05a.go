package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	type Vector struct {
		x1, y1, x2, y2 int
	}

	var vectors []Vector

	// Open file and read as words without whitespace
	ventsfile, err := os.Open("./vents.txt")
	chkerr(err)
	defer ventsfile.Close()

	ventsscan := bufio.NewScanner(ventsfile)
	ventsscan.Split(bufio.ScanWords)

	var tempvector Vector

	i := 1
	for ventsscan.Scan() {

		section := ventsscan.Text()

		if section == "->" {
			continue
		}

		if i == 1 {
			tempvectorslice := strings.Split(section, ",")
			splitvector1, err := strconv.Atoi(tempvectorslice[0])
			chkerr(err)
			splitvector2, err := strconv.Atoi(tempvectorslice[1])
			chkerr(err)
			tempvector.x1 = splitvector1
			tempvector.y1 = splitvector2
			i = 2
			continue
		}

		if i == 2 {
			tempvectorslice := strings.Split(section, ",")
			splitvector1, err := strconv.Atoi(tempvectorslice[0])
			chkerr(err)
			splitvector2, err := strconv.Atoi(tempvectorslice[1])
			chkerr(err)
			tempvector.x2 = splitvector1
			tempvector.y2 = splitvector2
			i = 1
		}

		vectors = append(vectors, tempvector)
	}

	intersects := make(map[int]map[int]int)
	var overlap int

	for _, vector := range vectors {
		fmt.Println(vector)
		//Check either x or y coords are equal ie. straight line

		var x, xa, xb, y, ya, yb int

		if vector.x1 != vector.x2 && vector.y1 != vector.y2 {
			continue
		} else if vector.x1 == vector.x2 {
			x = vector.x1
			// Make sure we are incrementing from lower y to greater y
			if vector.y2 < vector.y1 {
				ya = vector.y2
				yb = vector.y1
			} else {
				ya = vector.y1
				yb = vector.y2
			}

			for y := ya; y < yb; y++ {
				_, xok := intersects[x]
				if xok {
					_, yok := intersects[x][y]
					if yok {
						intersects[x][y]++
						overlap++
					} else {
						intersects[x][y] = 1
					}
				} else {
					intersects[x] = map[int]int{}
					intersects[x][y] = 1
				}
			}
		} else if vector.y1 == vector.y2 {
			y = vector.y1
			// Make sure we are incrementing from lower x to greater x
			if vector.x2 < vector.x1 {
				xa = vector.x2
				xb = vector.x1
			} else {
				xa = vector.x1
				xb = vector.x2
			}

			for x := xa; x < xb; x++ {
				_, xok := intersects[x]
				if xok {
					_, yok := intersects[x][y]
					if yok {
						intersects[x][y]++
						overlap++
					} else {
						intersects[x][y] = 1
					}
				} else {
					intersects[x] = map[int]int{}
					intersects[x][y] = 1
				}
			}
		}
	}

	fmt.Println(overlap)

}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
