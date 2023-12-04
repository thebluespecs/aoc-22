package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	readfile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	total := 0
	for fileScanner.Scan() {
		items1 := make(map[rune]bool)
		for _, item := range fileScanner.Text() {
			items1[item] = true
		}
		fileScanner.Scan()
		items2 := make(map[rune]bool)
		for _, item := range fileScanner.Text() {
			items2[item] = true
		}
		fileScanner.Scan()
		items3 := make(map[rune]bool)
		for _, item := range fileScanner.Text() {
			items3[item] = true
		}
		for item := range items1 {
			if items2[item] && items3[item] {
				i := strings.Index(Alphabets, string(item))
				total += i + 1
				break
			}
		}
	}
	fmt.Println(total)
	readfile.Close()
}
