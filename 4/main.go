package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readfile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
    total := 0
	for fileScanner.Scan() {
        ranges := strings.Split(fileScanner.Text(), ",")
        var lower, upper []int
        for _, r := range ranges {
            nums := strings.Split(r, "-")
            lower_int, _ := strconv.ParseInt(nums[0], 10, 32)
            upper_int, _ := strconv.ParseInt(nums[1], 10, 32)
            lower = append(lower, int(lower_int))
            upper = append(upper, int(upper_int))
        }
        // part 1
        // if (upper[0] >= upper[1] && lower[0] <= lower[1]) || (upper[1] >= upper[0] && lower[1] <= lower[0]) {
        // part 2
        if (lower[0] <= lower[1] && upper[0] >= lower[1]) || (lower[1] <= lower[0] && upper[1] >= lower[0]) {
            total += 1
        }
	}
    fmt.Println(total)
	readfile.Close()
}
