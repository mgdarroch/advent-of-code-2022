package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("day3/input_d3.txt")
	if err != nil {
		log.Fatal(err)
	}

	priorities := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "y": 24, "x": 25, "z": 26}
	fmt.Println(priorities["a"])
	defer f.Close()

	scanner := bufio.NewScanner(f)

	//var sum int = 0
	//
	//for scanner.Scan() {
	//	strVal := scanner.Text()
	//	mid := len(strVal) / 2
	//	pack1 := make(map[string]int)
	//	pairs := make(map[string]int)
	//	for i, r := range strVal {
	//		if i == 0 {
	//			fmt.Println("One:", strVal[:mid], "Two:", strVal[mid:])
	//		}
	//		if i < mid {
	//			pack1[string(r)] = pack1[string(r)] + 1
	//		} else {
	//			for k, _ := range pack1 {
	//				if k == string(r) {
	//					if pairs[string(r)] == 0 {
	//						pairs[string(r)] = pack1[string(r)] + 1
	//					} else {
	//						pairs[string(r)] = pairs[string(r)] + 1
	//					}
	//					fmt.Println("Pair found:", string(r), "Count:", pairs[string(r)])
	//				}
	//			}
	//		}
	//
	//	}
	//	fmt.Println("Pairs", pairs)
	//	fmt.Println(len(pack1), pack1)
	//
	//	for k, _ := range pairs {
	//		tmp := k
	//		num := 0
	//		if strings.ToUpper(tmp) == k {
	//			num = priorities[strings.ToLower(tmp)] + 26
	//		} else {
	//			num = priorities[strings.ToLower(tmp)]
	//		}
	//		sum += num
	//		fmt.Println("Priority Sum of", k, "equals", num)
	//	}
	//}
	//
	//fmt.Println("Final Sum:", sum)

	var count int = 1
	var groupCounter int = 1
	packSet := make(map[string]int)
	var sum int = 0
	for scanner.Scan() {
		strVal := scanner.Text()
		for _, r := range strVal {
			if count == 1 {
				packSet[string(r)] = 1
			}

			if count == 2 && packSet[string(r)] == 1 {
				packSet[string(r)] = count
			}

			if count == 3 && packSet[string(r)] == 2 {
				packSet[string(r)] = count
			}
		}

		if count == 3 {
			for k, v := range packSet {
				if v == 3 {
					tmp := strings.ToUpper(k)
					if k == tmp {
						sum += priorities[strings.ToLower(k)] + 26
					} else {
						sum += priorities[k]
					}

				}
			}
		}

		if count < 3 {
			count++
		} else if count == 3 {
			count = 1
			groupCounter++
			packSet = make(map[string]int)
		}
	}

	fmt.Println(sum)
}
