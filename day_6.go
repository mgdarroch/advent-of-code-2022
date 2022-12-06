package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input_d6.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := scanner.Scan()
	str := scanner.Text()
	fmt.Println(line, str)

	//var markerCount int = 0
	var positionCount int = 0

	for i, _ := range str {
		charSet := make(map[string]int)
		key := str[i : i+14]
		fmt.Println("Key:", key)
		for _, c := range key {
			if charSet[string(c)] == 0 {
				charSet[string(c)] = charSet[string(c)] + 1
			} else {
				break
			}
		}
		if len(charSet) == 14 {
			positionCount = i + 14
			break
		} else {
			continue
		}
	}

	fmt.Println("Number of Chars:", positionCount)
}
