package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items          []int64
	monkeyId       int
	itemsInspected int
}

func main() {
	monkeyMap := parseInput()
	// can only run one at a time because i committed to using pointers and actually editing the monkeys
	// probably a smarter way to do it.
	//throwItems(monkeyMap, true, 20)
	throwItems(monkeyMap, false, 10000)

	fmt.Println("Monkey Business:", getMonkeyBusiness(monkeyMap))
}

func printMonkeyItems(monkeyMap map[int]*Monkey) {
	for i := 0; i < len(monkeyMap); i++ {
		monkey := monkeyMap[i]
		fmt.Println("Monkey:", monkey.monkeyId, "Items:", monkey.items)
	}
}

func throwItems(monkeyMap map[int]*Monkey, relief bool, rounds int) {
	// in other words, prime number magic
	// then use the remainder as the new divisible
	// maffs
	x := int64(1)
	x *= 7
	x *= 19
	x *= 5
	x *= 11
	x *= 17
	x *= 13
	x *= 2
	x *= 3
	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeyMap); i++ {
			monkey := monkeyMap[i]
			for _, item := range monkey.items {
				if monkey.monkeyId == 0 {
					testVal := item * 19
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(7) == 0 {
						monkeyMap[2].items = append(monkeyMap[2].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[3].items = append(monkeyMap[3].items, testVal)
						monkey.items = monkey.items[1:]
					}
				} else if monkey.monkeyId == 1 {
					testVal := item + 1
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(19) == 0 {
						monkeyMap[4].items = append(monkeyMap[4].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[6].items = append(monkeyMap[6].items, testVal)
						monkey.items = monkey.items[1:]
					}
				} else if monkey.monkeyId == 2 {
					testVal := item + 6
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(5) == 0 {
						monkeyMap[7].items = append(monkeyMap[7].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[5].items = append(monkeyMap[5].items, testVal)
						monkey.items = monkey.items[1:]
					}
				} else if monkey.monkeyId == 3 {
					testVal := item + 5
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(11) == 0 {
						monkeyMap[5].items = append(monkeyMap[5].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[2].items = append(monkeyMap[2].items, testVal)
						monkey.items = monkey.items[1:]
					}
				} else if monkey.monkeyId == 4 {
					testVal := item * item
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(17) == 0 {
						monkeyMap[0].items = append(monkeyMap[0].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[3].items = append(monkeyMap[3].items, testVal)
						monkey.items = monkey.items[1:]
					}
				} else if monkey.monkeyId == 5 {
					testVal := item + 7
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(13) == 0 {
						monkeyMap[1].items = append(monkeyMap[1].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[7].items = append(monkeyMap[7].items, testVal)
						monkey.items = monkey.items[1:]
					}
				} else if monkey.monkeyId == 6 {
					testVal := item * 7
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(2) == 0 {
						monkeyMap[0].items = append(monkeyMap[0].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[4].items = append(monkeyMap[4].items, testVal)
						monkey.items = monkey.items[1:]
					}
				} else if monkey.monkeyId == 7 {
					testVal := item + 2
					if relief {
						testVal = int64(int(math.Floor(float64(testVal / 3))))
					}
					testVal = testVal % int64(x)
					monkey.itemsInspected = monkey.itemsInspected + 1
					if testVal%int64(3) == 0 {
						monkeyMap[1].items = append(monkeyMap[1].items, testVal)
						monkey.items = monkey.items[1:]
					} else {
						monkeyMap[6].items = append(monkeyMap[6].items, testVal)
						monkey.items = monkey.items[1:]
					}
				}
			}
		}
	}
}

func parseInput() map[int]*Monkey {
	f, err := os.Open("day11/input_d11.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	monkeyMap := make(map[int]*Monkey)
	monkeCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strVal := scanner.Text()
		strVal = strings.TrimSpace(strVal)
		strVal = strings.ReplaceAll(strVal, ",", "")
		args := strings.Split(strVal, " ")
		if args[0] == "Monkey" {
			numStr := args[1]
			num, err := strconv.Atoi(numStr[:1])
			monkeCount = num
			if err == nil {
				items := make([]int64, 0)
				monke := Monkey{items: items, monkeyId: num, itemsInspected: 0}
				monkeyMap[num] = &monke
			}
		}
		if args[0] == "Starting" {
			for _, v := range args[2:] {
				num, _ := strconv.Atoi(v)

				monkeyMap[monkeCount].items = append(monkeyMap[monkeCount].items, int64(num))
			}
		}
	}
	return monkeyMap
}

func getMonkeyBusiness(monkeyMap map[int]*Monkey) int {
	itemsInspected := make([]int, len(monkeyMap))
	for i, v := range monkeyMap {
		itemsInspected[i] = v.itemsInspected
	}

	sort.Ints(itemsInspected)

	return itemsInspected[len(itemsInspected)-2] * itemsInspected[len(itemsInspected)-1]
}
