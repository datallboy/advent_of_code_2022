package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	var totalSumOfPrority int

	for fileScan.Scan() {
		var line []byte = fileScan.Bytes()
		var halfLength int = len(line) / 2

		var firstSack []byte = line[0:halfLength]
		var secondSack []byte = line[halfLength:]

		var alreadyFound []byte

		// Loop through firstSack and compare each byte to the bytes in secondSack
		for _, firstSackElement := range firstSack {
			for _, secondSackElement := range secondSack {

				if contains(firstSackElement, alreadyFound) {
					break
				}

				if firstSackElement == secondSackElement {
					// 91 = ASCII, uppercase letters
					if firstSackElement < 91 {
						totalSumOfPrority += int(firstSackElement - 38)
					} else {
						totalSumOfPrority += int(firstSackElement - 96)
					}
					alreadyFound = append(alreadyFound, firstSackElement)
				}
			}
		}
	}

	fmt.Println("Total Sum of Priority: ", totalSumOfPrority)
}

func contains(b byte, byteArray []byte) bool {
	for _, v := range byteArray {
		if v == b {
			return true
		}
	}

	return false
}
