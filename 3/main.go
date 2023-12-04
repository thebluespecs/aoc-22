package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var Alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func name() {
    readfile, err := os.Open("./input.txt")
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readfile)
    fileScanner.Split(bufio.ScanLines)
    total := 0 
    for fileScanner.Scan() {
        items := make(map[rune]bool)
        for _, leftItem := range fileScanner.Text()[0:len(fileScanner.Text())/2] {
            items[leftItem] = true
        }
        for _, rightItem := range fileScanner.Text()[len(fileScanner.Text())/2:] {
            if items[rightItem] {
                //get the index of the item in the alphabets string
                i := strings.Index(Alphabets, string(rightItem))
                total += i+1
                break;
            }
        } 
    }
    fmt.Println(total)
    readfile.Close()
}
