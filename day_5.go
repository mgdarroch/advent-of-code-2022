package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func NewStack() *Stack {
	return &Stack{}
}

func Reverse(input []string) []string {
	inputLen := len(input)
	output := make([]string, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}

func main() {

	f, err := os.Open("input_d5.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var line int = 1

	var stack1 []string = make([]string, 0)
	var stack2 []string = make([]string, 0)
	var stack3 []string = make([]string, 0)
	var stack4 []string = make([]string, 0)
	var stack5 []string = make([]string, 0)
	var stack6 []string = make([]string, 0)
	var stack7 []string = make([]string, 0)
	var stack8 []string = make([]string, 0)
	var stack9 []string = make([]string, 0)

	crateMap := map[int][]string{1: stack1, 2: stack2, 3: stack3, 4: stack4, 5: stack5, 6: stack6, 7: stack7, 8: stack8, 9: stack9}
	fmt.Println(crateMap)

	for scanner.Scan() {
		if line == 9 {
			break
		}
		strVal := scanner.Text()
		charCount := 1
		stackCount := 0
		for _, c := range strVal {
			if charCount == 4 {
				charCount = 1
				continue
			}
			if charCount == 2 {
				if string(c) == " " {
					charCount++
					stackCount++
					continue
				}
				stackCount++
				crateMap[stackCount] = append(crateMap[stackCount], string(c))
			}
			charCount++
		}
		line++
	}

	for k, _ := range crateMap {
		crateMap[k] = Reverse(crateMap[k])
	}
	fmt.Println("Crate Map:", crateMap)
	mapTotal := 0
	for k, _ := range crateMap {
		mapTotal += len(crateMap[k])
	}
	fmt.Println("Map Total:", mapTotal)

	line = 0
	//for scanner.Scan() {
	//	strVal := scanner.Text()
	//	if strVal == "" {
	//		continue
	//	}
	//	valSlice := strings.Split(strVal, " ")
	//	moves, _ := strconv.Atoi(valSlice[1])
	//	crateFrom, _ := strconv.Atoi(valSlice[3])
	//	crateTo, _ := strconv.Atoi(valSlice[5])
	//	for i := 1; i <= moves; i++ {
	//		fmt.Println("Moving from Crate", crateFrom, crateMap[crateFrom], "to Crate", crateTo, crateMap[crateTo])
	//		n := len(crateMap[crateFrom]) - 1
	//		if n < 0 {
	//			fmt.Println("length is zero")
	//			continue
	//		} else {
	//			topString := crateMap[crateFrom][n]
	//			crateMap[crateTo] = append(crateMap[crateTo], topString)
	//			crateMap[crateFrom] = crateMap[crateFrom][:n]
	//			fmt.Println("Moved", topString, "from", crateMap[crateFrom], "to", crateMap[crateTo])
	//		}
	//	}
	//}

	for scanner.Scan() {
		strVal := scanner.Text()
		if strVal == "" {
			continue
		}
		valSlice := strings.Split(strVal, " ")
		moves, _ := strconv.Atoi(valSlice[1])
		crateFrom, _ := strconv.Atoi(valSlice[3])
		crateTo, _ := strconv.Atoi(valSlice[5])
		n := 0
		fmt.Println("Moving", moves, "from Crate", crateFrom, crateMap[crateFrom], "to Crate", crateTo, crateMap[crateTo])
		if (len(crateMap[crateFrom]) - moves) <= 0 {
			fmt.Println("Move is the entire slice")
			n = len(crateMap[crateFrom])
			topSlice := crateMap[crateFrom][0:]
			fmt.Println("Slice:", topSlice)
			for _, s := range topSlice {
				crateMap[crateTo] = append(crateMap[crateTo], s)
				crateMap[crateFrom] = crateMap[crateFrom][:0]
				fmt.Println("Moved", s, "from", crateMap[crateFrom], "to", crateMap[crateTo])
			}
		} else {
			fmt.Println("Moving", n)
			n = len(crateMap[crateFrom]) - moves
			topSlice := crateMap[crateFrom][n:]
			fmt.Println("Slice:", topSlice)
			for _, s := range topSlice {
				crateMap[crateTo] = append(crateMap[crateTo], s)
				crateMap[crateFrom] = crateMap[crateFrom][:n]
				fmt.Println("Moved", s, "from", crateMap[crateFrom], "to", crateMap[crateTo])
			}
		}
	}

	fmt.Println(crateMap)
	mapTotal = 0
	for k, _ := range crateMap {
		mapTotal += len(crateMap[k])
	}
	fmt.Println("Map Total:", mapTotal)
}
