package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	// Open file and read as words without whitespace
	signalfile, err := os.Open("./signals.txt")
	chkerr(err)
	defer signalfile.Close()

	// Slice of the 14 codes in original character order

	signalscan := bufio.NewScanner(signalfile)
	signalscan.Split(bufio.ScanLines)
	var ctotal int

	for signalscan.Scan() {

		var signalslc []string
		signalordslc := make([]string, 14, 14)

		signalstr := signalscan.Text()
		signalslcbar := strings.Split(signalstr, " ")
		signalslc = append(signalslcbar[:10], signalslcbar[11:]...)

		for i, s := range signalslc {
			signalordslc[i] = SortStringByCharacter(s)
		}

		nums := make(map[int]string)
		codes := make(map[string]int)

		for _, s := range signalordslc {
			switch len(s) {
			case 2:
				nums[1] = s
				codes[s] = 1
			case 3:
				nums[7] = s
				codes[s] = 7
			case 4:
				nums[4] = s
				codes[s] = 4
			case 7:
				nums[8] = s
				codes[s] = 8
			}
		}

		// If 1 exists 5 sided figure including both segments in 1 is 3

		if l, ok := nums[1]; ok {
			for _, t := range signalordslc {
				c := 0
				if len(t) == 5 {
					for _, s := range l {
						for _, u := range t {
							if s == u {
								c++
							}
						}
					}
					if c == 2 {
						nums[3] = t
						codes[t] = 3
					}
				}
			}
		}

		// If 1 exists 6 sided figure NOT including both segments in 1 is 6

		if l, ok := nums[1]; ok {
			for _, t := range signalordslc {
				c := 0
				if len(t) == 6 {
					for _, s := range l {
						for _, u := range t {
							if s == u {
								c++
							}
						}
					}
					if c != 2 {
						nums[6] = t
						codes[t] = 6
					}
				}
			}
		}
		// If 7 exists 5 sided figure including all segments in 7 is 3

		if l, ok := nums[7]; ok {
			for _, t := range signalordslc {
				c := 0
				if len(t) == 5 {
					for _, s := range l {
						for _, u := range t {
							if s == u {
								c++
							}
						}
					}
					if c == 3 {
						nums[3] = t
						codes[t] = 3
					}
				}
			}
		}

		// If 7 exists 6 sided figure NOT including both segments in 7 is 6

		if l, ok := nums[7]; ok {
			for _, t := range signalordslc {
				c := 0
				if len(t) == 6 {
					for _, s := range l {
						for _, u := range t {
							if s == u {
								c++
							}
						}
					}
					if c != 3 {
						nums[6] = t
						codes[t] = 6
					}
				}
			}
		}

		// If 4 exists 5 sided figure including two segments in 4 is 2

		if l, ok := nums[4]; ok {
			for _, t := range signalordslc {
				c := 0
				if len(t) == 5 {
					for _, s := range l {
						for _, u := range t {
							if s == u {
								c++
							}
						}
					}
					if c == 2 {
						nums[2] = t
						codes[t] = 2
					}
				}
			}
		}

		// If 4 exists 6 sided figure including all segments in 4 is 9

		if l, ok := nums[4]; ok {
			for _, t := range signalordslc {
				c := 0
				if len(t) == 6 {
					for _, s := range l {
						for _, u := range t {
							if s == u {
								c++
							}
						}
					}
					if c == 4 {
						nums[9] = t
						codes[t] = 9
					}
				}
			}
		}

		_, sok := nums[6]
		_, nok := nums[9]
		if sok && nok {
			for _, t := range signalordslc {
				if len(t) == 6 && t != nums[6] && t != nums[9] {
					nums[0] = t
					codes[t] = 0
				}
			}
		}

		_, tok := nums[2]
		_, hok := nums[3]
		if tok && hok {
			for _, t := range signalordslc {
				if len(t) == 5 && t != nums[2] && t != nums[3] {
					nums[5] = t
					codes[t] = 5
				}
			}
		}

		thousands := codes[signalordslc[10]]
		hundreds := codes[signalordslc[11]]
		tens := codes[signalordslc[12]]
		units := codes[signalordslc[13]]

		total := (thousands * 1000) + (hundreds * 100) + (tens * 10) + units
		ctotal = ctotal + total

		fmt.Println(signalordslc)
		fmt.Println(nums)
		fmt.Println(codes)
		fmt.Println(total)
	}

	fmt.Println("")
	fmt.Println(ctotal)

}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
