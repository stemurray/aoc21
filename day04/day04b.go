package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	type Card struct {
		row     [5]int
		col     [5]int
		total   int
		checked int
	}

	type Cell struct {
		board int
		row   int
		col   int
	}

	// A map which uses the ball # as an index to which Cells on which Cards have that #
	cells := make(map[int][]Cell)

	// A slice which will hold the Cards for the game
	game := make([]Card, 1, 150)

	// Open file and read as words without whitespace
	bingofile, err := os.Open("./bingo.txt")
	chkerr(err)
	defer bingofile.Close()

	bingoscan := bufio.NewScanner(bingofile)
	bingoscan.Split(bufio.ScanWords)

	// Read first line (which is comma separated) and add to a slice of strings
	bingoscan.Scan()
	ballstr := bingoscan.Text()
	balls := strings.Split(ballstr, ",")

	// Populate the map of cells on cards and calculate total for each card
	var cardnum int
	var rownum int
	var colnum int
	var cardtotal int
	newcard := Card{}
	var numcards int

	for bingoscan.Scan() {
		c := Cell{cardnum, rownum, colnum}
		ballnum, err := strconv.Atoi(bingoscan.Text())
		chkerr(err)
		cells[ballnum] = append(cells[ballnum], c)
		cardtotal = cardtotal + ballnum

		if colnum == 4 && rownum == 4 {
			colnum = 0
			rownum = 0
			game[cardnum].total = cardtotal
			cardtotal = 0
			cardnum++
			game = append(game, newcard)
			numcards++
		} else if colnum == 4 && rownum < 4 {
			colnum = 0
			rownum++
		} else {
			colnum++
		}
	}

	cardcomplete := make([]int, numcards+1)

	// FOR EACH BALL
	for _, b := range balls {

		bint, err := strconv.Atoi(b)
		chkerr(err)

		// FOR EACH MAP WHERE KEY IS BALL #
		for _, cl := range cells[bint] {

			// IF THE CARD ROW REFEERENCED BY MAP CELL == 4, THIS ONE WOULD MAKE IT WIN
			if game[cl.board].row[cl.row] == 4 {
				if cardcomplete[cl.board] == 0 {
					cardcomplete[cl.board] = 1
					numcards--
				}
			} else {
				// IF NOT WINNER INCREMENT THE ROW VALUE OF THIS CARD
				game[cl.board].row[cl.row]++
			}

			// IF THE CARD COL REFEERENCED BY MAP CELL == 4, THIS ONE WOULD MAKE IT WIN
			if game[cl.board].col[cl.col] == 4 {
				if cardcomplete[cl.board] == 0 {
					cardcomplete[cl.board] = 1
					numcards--
				}
			} else {
				// IF NOT WINNER INCREMENT THE COL VALUE OF THIS CARD
				game[cl.board].col[cl.col]++
			}

			// ADD THE BALL # TO THE CARD BALL TOTAL
			game[cl.board].checked = game[cl.board].checked + bint

			if numcards == 0 {
				fmt.Println((game[cl.board].total - game[cl.board].checked) * bint)
				break
			}
		}

		if numcards == 0 {
			break
		}
	}
}

//Functions

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
