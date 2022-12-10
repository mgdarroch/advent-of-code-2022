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

	// noop begins execution.  x = 1 during the first cycle.  after the first cycle, noop finishes execution doing nothing
	// addx -5 = add -5 to x
	// signal strength = the cycle number multiplied by the value of the x register
	// addx V takes two cycles to complete.
	// noop takes one cycle.
	// find signal strength during 20th, 60th, 100th, 140th, 180th and 220th cycles.  sum them.

	f, err := os.Open("day10/input_d10.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	cycle := 0
	x := 1
	signalSum := 0

	screen := make([][]string, 6)

	for i, _ := range screen {
		for j := 0; j < 40; j++ {
			screen[i] = append(screen[i], ".")
		}
	}

	printScreen(screen)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		switch instruction := args[0]; instruction {
		case "noop":
			drawPixel(screen, cycle, x)
			cycle += 1
			switch cycle {
			case 20, 60, 100, 140, 180, 220:
				signalSum += cycle * x
			}
			//fmt.Println("Cycle:", cycle, "X-val:", x, "Signal Strength:", cycle*x)
		case "addx":
			num, _ := strconv.Atoi(args[1])
			for i := 0; i < 2; i++ {
				drawPixel(screen, cycle, x)
				if i == 0 {
					cycle++
					switch cycle {
					case 20, 60, 100, 140, 180, 220:
						signalSum += cycle * x
					}
					//fmt.Println("Cycle:", cycle, "X-val:", x, "Signal Strength:", cycle*x)
				}
				if i == 1 {
					cycle++
					switch cycle {
					case 20, 60, 100, 140, 180, 220:
						signalSum += cycle * x
					}
					//fmt.Println("Cycle:", cycle, "X-val:", x, "Signal Strength:", cycle*x)
					x += num
				}
			}
		}
	}
	fmt.Println("SignalSum:", signalSum)
}

func drawPixel(screen [][]string, cycle int, register int) {
	fmt.Println("Cycle:", cycle)
	fmt.Println("Register:", register)

	registerInsertMapping := register
	spriteMaxInsert := registerInsertMapping + 1
	spriteMinInsert := registerInsertMapping - 1

	draw := false

	if cycle < 40 {
		if cycle == register || cycle == register-1 || cycle == register+1 {
			draw = true
		}

		if !draw {
			return
		}

		if cycle == register-1 && spriteMinInsert >= 0 {
			screen[0][spriteMinInsert] = "#"
		}
		if cycle == register && register >= 0 {
			screen[0][registerInsertMapping] = "#"
		}
		if cycle == register+1 && spriteMaxInsert < 40 {
			screen[0][spriteMaxInsert] = "#"
		}
	} else if cycle < 80 {
		cycle = cycle - 40
		if cycle == register || cycle == register-1 || cycle == register+1 {
			draw = true
		}

		if !draw {
			return
		}

		if cycle == register-1 && spriteMinInsert >= 0 {
			screen[1][spriteMinInsert] = "#"
		}
		if cycle == register && register >= 0 {
			screen[1][registerInsertMapping] = "#"
		}
		if cycle == register+1 && spriteMaxInsert < 40 {
			screen[1][spriteMaxInsert] = "#"
		}
	} else if cycle < 120 {
		cycle = cycle - 80
		if cycle == register || cycle == register-1 || cycle == register+1 {
			draw = true
		}

		if !draw {
			return
		}

		if cycle == register-1 && spriteMinInsert >= 0 {
			screen[2][spriteMinInsert] = "#"
		}
		if cycle == register && register >= 0 {
			screen[2][registerInsertMapping] = "#"
		}
		if cycle == register+1 && spriteMaxInsert < 40 {
			screen[2][spriteMaxInsert] = "#"
		}
	} else if cycle < 160 {
		cycle = cycle - 120
		if cycle == register || cycle == register-1 || cycle == register+1 {
			draw = true
		}

		if !draw {
			return
		}

		if cycle == register-1 && spriteMinInsert >= 0 {
			screen[3][spriteMinInsert] = "#"
		}
		if cycle == register && register >= 0 {
			screen[3][registerInsertMapping] = "#"
		}
		if cycle == register+1 && spriteMaxInsert < 40 {
			screen[3][spriteMaxInsert] = "#"
		}
	} else if cycle < 200 {
		cycle = cycle - 160
		if cycle == register || cycle == register-1 || cycle == register+1 {
			draw = true
		}

		if !draw {
			return
		}

		if cycle == register-1 && spriteMinInsert >= 0 {
			screen[4][spriteMinInsert] = "#"
		}
		if cycle == register && register >= 0 {
			screen[4][registerInsertMapping] = "#"
		}
		if cycle == register+1 && spriteMaxInsert < 40 {
			screen[4][spriteMaxInsert] = "#"
		}
	} else if cycle < 240 {
		cycle = cycle - 200
		if cycle == register || cycle == register-1 || cycle == register+1 {
			draw = true
		}

		if !draw {
			return
		}

		if cycle == register-1 && spriteMinInsert >= 0 {
			screen[5][spriteMinInsert] = "#"
		}
		if cycle == register && register >= 0 {
			screen[5][registerInsertMapping] = "#"
		}
		if cycle == register+1 && spriteMaxInsert < 40 {
			screen[5][spriteMaxInsert] = "#"
		}
	}

	printScreen(screen)
}

func printScreen(screen [][]string) {
	for _, v := range screen {
		fmt.Println(v)
	}
}
