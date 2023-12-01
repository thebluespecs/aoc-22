package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	readfile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
    highest, secondHighest, thirdHighest := 0, 0, 0
    current := 0
	for fileScanner.Scan() {
        if fileScanner.Text() == "" {
            if current > highest {
                thirdHighest = secondHighest
                secondHighest = highest
                highest = current
            } else if current > secondHighest {
                thirdHighest = secondHighest
                secondHighest = current
            } else if current > thirdHighest {
                thirdHighest = current
            }
            current = 0
        } else {
            weight, err := strconv.Atoi(fileScanner.Text())
            if err != nil {
                fmt.Println(err)
            }
            current += weight
        }
	}
	readfile.Close()
    fmt.Println(highest, secondHighest, thirdHighest)
    fmt.Println(highest + secondHighest + thirdHighest)
}
