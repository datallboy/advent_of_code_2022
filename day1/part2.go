package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	var fileLines []int
	var calories_per_elf int

	for fileScan.Scan() {
		calorie, _ := strconv.Atoi(fileScan.Text())
		calories_per_elf += calorie

		if fileScan.Text() == "" {
			fileLines = append(fileLines, calories_per_elf)
			calories_per_elf = 0
		}
	}

	sort.Ints(fileLines)
	fmt.Println("Calories of Top Elf: ", fileLines[len(fileLines)-1])

	sum_of_top_three_elves := fileLines[len(fileLines)-1] + fileLines[len(fileLines)-2] + fileLines[len(fileLines)-3]
	fmt.Println("Sum of Calories for Top Three Elves: ", sum_of_top_three_elves)
}
