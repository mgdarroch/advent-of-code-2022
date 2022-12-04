package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input_d4.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var count int = 0
	var lowSetTracker int = 0
	var highSetTracker int = 0

	//for scanner.Scan() {
	//	strVal := scanner.Text()
	//	stringSlice := strings.Split(strVal, ",")
	//	set1 := strings.Split(stringSlice[0], "-")
	//	set2 := strings.Split(stringSlice[1], "-")
	//	set1num1, _ := strconv.Atoi(set1[0])
	//	set1num2, _ := strconv.Atoi(set1[1])
	//	set2num1, _ := strconv.Atoi(set2[0])
	//	set2num2, _ := strconv.Atoi(set2[1])
	//
	//	min := 0
	//	if set1num1 < set2num1 {
	//		min, err = strconv.Atoi(set1[0])
	//		lowSetTracker = 1
	//	} else if set2num1 < set1num1 {
	//		min, err = strconv.Atoi(set2[0])
	//		lowSetTracker = 2
	//	} else {
	//		min, err = strconv.Atoi(set1[0])
	//		lowSetTracker = 3
	//	}
	//
	//	max := 0
	//	if set1num2 > set2num2 {
	//		max, err = strconv.Atoi(set1[1])
	//		highSetTracker = 1
	//	} else if set2num2 > set1num2 {
	//		max, err = strconv.Atoi(set2[1])
	//		highSetTracker = 2
	//	} else {
	//		max, err = strconv.Atoi(set1[1])
	//		highSetTracker = 3
	//	}
	//
	//	if lowSetTracker != 3 && highSetTracker != 3 {
	//		if lowSetTracker == highSetTracker {
	//			count++
	//		}
	//	} else {
	//		count++
	//	}
	//
	//	fmt.Println("Low Set:", lowSetTracker, "High Set:", highSetTracker, "Min:", min, "Max:", max, "Count:", count, "Set 1:", set1, "Set 2:", set2)
	//	//fmt.Println("Min:", min, "Max:", max)
	//}

	for scanner.Scan() {
		strVal := scanner.Text()
		stringSlice := strings.Split(strVal, ",")
		set1 := strings.Split(stringSlice[0], "-")
		set2 := strings.Split(stringSlice[1], "-")
		set1num1, _ := strconv.Atoi(set1[0])
		set1num2, _ := strconv.Atoi(set1[1])
		set2num1, _ := strconv.Atoi(set2[0])
		set2num2, _ := strconv.Atoi(set2[1])

		min := 0
		if set1num1 < set2num1 {
			min, err = strconv.Atoi(set1[0])
			lowSetTracker = 1
		} else if set2num1 < set1num1 {
			min, err = strconv.Atoi(set2[0])
			lowSetTracker = 2
		} else {
			min, err = strconv.Atoi(set1[0])
			lowSetTracker = 3
		}

		max := 0
		if set1num2 > set2num2 {
			max, err = strconv.Atoi(set1[1])
			highSetTracker = 1
		} else if set2num2 > set1num2 {
			max, err = strconv.Atoi(set2[1])
			highSetTracker = 2
		} else {
			max, err = strconv.Atoi(set1[1])
			highSetTracker = 3
		}

		if lowSetTracker != highSetTracker && lowSetTracker != 3 && highSetTracker != 3 {
			if lowSetTracker == 1 {
				if set2num1 < set1num2 {
					count++
				} else if set1num2 >= set2num1 {
					count++
				}
			} else {
				if set1num1 < set2num2 {
					count++
				} else if set2num2 >= set1num1 {
					count++
				}
			}
		} else if lowSetTracker != 3 && highSetTracker != 3 {
			if lowSetTracker == highSetTracker {
				count++
			}
		} else {
			count++
		}

		fmt.Println("Low Set:", lowSetTracker, "High Set:", highSetTracker, "Min:", min, "Max:", max, "Count:", count, "Set 1:", set1, "Set 2:", set2)
		//fmt.Println("Min:", min, "Max:", max)
	}
	fmt.Println("Final Count:", count)

}
