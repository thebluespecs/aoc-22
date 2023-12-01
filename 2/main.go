package main

import (
	"bufio"
	"fmt"
	"os"
)

var table = map[string]int{
    "A X": 4,
    "A Y": 8,
    "A Z": 3,
    "B X": 1,
    "B Y": 5,
    "B Z": 9,
    "C X": 7,
    "C Y": 2,
    "C Z": 6,
}

var table2 = map[string]int{
    "A X": 3,
    "A Y": 4,
    "A Z": 8,
    "B X": 1,
    "B Y": 5,
    "B Z": 9,
    "C X": 2,
    "C Y": 6,
    "C Z": 7,
}

func main() {
	readfile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
    totalScore := 0
	for fileScanner.Scan() {
        switch fileScanner.Text() {
        case "A X":
            totalScore += table2["A X"]
        case "A Y":
            totalScore += table2["A Y"]
        case "A Z":
            totalScore += table2["A Z"]
        case "B X":
            totalScore += table2["B X"]
        case "B Y":
            totalScore += table2["B Y"]
        case "B Z":
            totalScore += table2["B Z"]
        case "C X":
            totalScore += table2["C X"]
        case "C Y":
            totalScore += table2["C Y"]
        case "C Z":
            totalScore += table2["C Z"]
        }
	}
	readfile.Close()
    fmt.Println(totalScore)
}
