package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day12/input_d12.txt")
	height, start, end := readMap(string(input))

	part1 := getShortestPath(height, map[Position]int{}, Visit{start, 0},
		func(dh int) bool { return dh <= 1 },
		func(p Position) bool { return end == p })

	part2 := getShortestPath(height, map[Position]int{}, Visit{end, 0},
		func(dh int) bool { return dh >= -1 },
		func(p Position) bool { return height[p] == 'a' })

	fmt.Printf("part1: %v\npart2: %v\n", part1, part2)
}

func getShortestPath(heightMap map[Position]int, visits map[Position]int, visitPosition Visit, okToMove func(d int) bool, target func(Position) bool) int {

	//Initialize a stack with the starting position loc and a min variable to keep track of the minimum distance found so far.
	//	While the stack is not empty:
	//Pop the top element from the stack and set it as the current position.
	//	If the current position is the target position, update the min variable if the current distance is smaller than the current value of min.
	//	If the current position has been visited before and the previous distance was smaller than or equal to the current distance, skip this position.
	//	Otherwise, mark the current position as visited with the current distance.
	//	For each of the four possible neighbors (up, down, left, and right):
	//If the neighbor position is inside the map and the height difference is allowed by the ok function, append the neighbor position to the stack.
	//	Return the final value of min as the result.

	min := 9999
	stack := []Visit{visitPosition}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if target(current.Position) {
			if current.distance < min {
				min = current.distance
			}
			continue
		} else if old, visited := visits[current.Position]; visited && old <= current.distance {
			continue
		}

		visits[current.Position] = current.distance

		for _, m := range []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := Visit{Position{current.x + m.x, current.y + m.y}, current.distance + 1}
			if h, inside := heightMap[next.Position]; inside {
				deltaH := h - heightMap[current.Position]
				if !okToMove(deltaH) {
					continue
				}
				stack = append(stack, next)
			}
		}
	}
	return min
}

func readMap(input string) (heightMap map[Position]int, start, end Position) {
	rows := strings.Split(input, "\n")
	heightMap = map[Position]int{}
	for y, row := range rows {
		for x, c := range row {
			if c == 'S' {
				start = Position{x, y}
				c = 'a'
			} else if c == 'E' {
				end = Position{x, y}
				c = 'z'
			}
			heightMap[Position{x, y}] = int(c)
		}
	}
	return heightMap, start, end
}

type Visit struct {
	Position
	distance int
}

type Position struct {
	x int
	y int
}
