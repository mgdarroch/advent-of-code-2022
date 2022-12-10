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

	f, err := os.Open("day1/input_d1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var count = 0
	var highVal = 0
	x := make([]int, 0)

	for scanner.Scan() {
		strVal := scanner.Text()
		if strVal == "" {
			if count > highVal {
				highVal = count
			}
			x = append(x, count)
			count = 0
			continue
		}

		intVal, err := strconv.Atoi(strVal)

		if err == nil {
			count += intVal
		}
	}

	sort.Sort(sort.IntSlice(x))
	y := x[237:]
	var sum int = 0

	for _, num := range y {
		sum += num
	}

	fmt.Println("Top 3 Sum: ", sum)

	fmt.Println("Highest Value is: ", highVal)
	fmt.Println(len(x))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
