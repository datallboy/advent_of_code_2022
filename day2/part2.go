package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 65 & 88 = A & X = Rock 		1 point
// 66 & 89 = B & Y = Paper		2 points
// 67 & 90 = C & Z = Scissors	3 points

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	opponentWinMap := map[byte]byte{
		65: 90,
		66: 88,
		67: 89,
	}

	opponentLoseMap := map[byte]byte{
		65: 89,
		66: 90,
		67: 88,
	}

	drawMap := map[byte]byte{
		65: 88,
		66: 89,
		67: 90,
	}

	decimalToPointsMap := map[byte]int{
		88: 1,
		89: 2,
		90: 3,
	}

	const DRAW_POINTS = 3
	const WIN_POINTS = 6

	var points int

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	for fileScan.Scan() {
		opponent := fileScan.Bytes()[0]
		you := fileScan.Bytes()[2]

		switch you {
		case byte('X'):
			// need to lose
			points += decimalToPointsMap[opponentWinMap[opponent]]
		case byte('Y'):
			// need to draw
			points += DRAW_POINTS + decimalToPointsMap[drawMap[opponent]]
		case byte('Z'):
			points += WIN_POINTS + decimalToPointsMap[opponentLoseMap[opponent]]
		}
	}

	fmt.Println("Total Points: ", points)
}
