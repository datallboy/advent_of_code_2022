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
		88: 66,
		89: 67,
		90: 65,
	}

	drawMap := map[byte]byte{
		88: 65,
		89: 66,
		90: 67,
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

		if opponent == drawMap[you] {
			// draw
			points += DRAW_POINTS + decimalToPointsMap[you]
		} else if opponent == opponentWinMap[you] {
			// you lost
			points += decimalToPointsMap[you]
		} else {
			// you won!
			points += WIN_POINTS + decimalToPointsMap[you]
		}
	}

	fmt.Println("Total Points: ", points)
}
