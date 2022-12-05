package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var calories_per_elf int
	var elf_with_most_calories int

	for fileScan.Scan() {
		calorie, _ := strconv.Atoi(fileScan.Text())
		calories_per_elf += calorie

		if fileScan.Text() == "" {
			if calories_per_elf > elf_with_most_calories {
				elf_with_most_calories = calories_per_elf
			}
			calories_per_elf = 0
		}
	}

	fmt.Printf("Most Calories: %d", elf_with_most_calories)
}
