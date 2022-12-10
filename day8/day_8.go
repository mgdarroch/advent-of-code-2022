package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getHighestScenicScore(grid [][]int) int {
	x := len(grid)
	upCount := 0
	downCount := 0
	leftCount := 0
	rightCount := 0
	maxScenicScore := 0

	for i := 0; i < x; i++ {
		for j := 0; j < len(grid[i]); j++ {
			curr := grid[i][j]
			upCount = getUpCount(grid, i, j, curr)
			downCount = getDownCount(grid, i, j, curr)
			leftCount = getLeftCount(grid, i, j, curr)
			rightCount = getRightCount(grid, i, j, curr)
			tmpScenicScore := upCount * downCount * leftCount * rightCount
			if tmpScenicScore > maxScenicScore {
				maxScenicScore = tmpScenicScore
			}
		}
	}
	return maxScenicScore
}

func getRightCount(grid [][]int, i int, j int, curr int) int {
	count := 0
	for j < len(grid[i])-1 {
		count++
		if grid[i][j+1] >= curr {
			break
		}
		j++
	}
	return count
}

func getLeftCount(grid [][]int, i int, j int, curr int) int {
	count := 0
	for j > 0 {
		count++
		if grid[i][j-1] >= curr {
			break
		}
		j--

	}
	return count
}

func getDownCount(grid [][]int, i int, j int, curr int) int {
	count := 0
	for i < len(grid)-1 {
		count++
		if grid[i+1][j] >= curr {
			break
		}
		i++
	}
	return count
}

func getUpCount(grid [][]int, i int, j int, curr int) int {
	count := 0
	for i > 0 {
		count++
		if grid[i-1][j] >= curr {
			break
		}
		i--
	}
	return count
}

func countVisibleTrees(grid [][]int) int {

	// Edges
	x := len(grid)
	count := 0
	count += len(grid[0])
	count += len(grid[x-1])
	y := len(grid[:x-2])
	count += y * 2

	// Interior
	for i := 1; i < x-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			curr := grid[i][j]
			if checkUp(grid, i, j, curr) {
				count++
			} else if checkDown(grid, i, j, curr) {
				count++
			} else if checkLeft(grid, i, j, curr) {
				count++
			} else if checkRight(grid, i, j, curr) {
				count++
			} else {
				continue
			}
		}
	}

	return count
}

func checkUp(grid [][]int, i, j int, curr int) bool {
	for i > 0 {
		if grid[i-1][j] >= curr {
			return false
		}
		i--
	}
	return true
}

func checkDown(grid [][]int, i, j int, curr int) bool {
	for i < len(grid)-1 {
		if grid[i+1][j] >= curr {
			return false
		}
		i++
	}
	return true
}

func checkLeft(grid [][]int, i, j int, curr int) bool {
	for j > 0 {
		if grid[i][j-1] >= curr {
			return false
		}
		j--
	}
	return true
}

func checkRight(grid [][]int, i, j int, curr int) bool {
	for j < len(grid[i])-1 {
		if grid[i][j+1] >= curr {
			return false
		}
		j++
	}
	return true
}

func main() {
	f, err := os.Open("day8/input_d8.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, ch := range line {
			n, _ := strconv.Atoi(string(ch))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	fmt.Println(countVisibleTrees(grid))
	fmt.Println(getHighestScenicScore(grid))
}
