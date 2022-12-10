package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// A = Rock, B = Paper, C = Scissors
	// Y = Paper, X = Rock, Z = Scissors
	f, err := os.Open("day2/input_d2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int = 0

	//for scanner.Scan() {
	//	strVal := scanner.Text()
	//	fmt.Println(strVal)
	//	switch outcome := strVal; outcome {
	//	case "A Y":
	//		fmt.Println("Rock loses to Paper.")
	//		total += 8
	//	case "A X":
	//		fmt.Println("Rock draws with Rock.")
	//		total += 4
	//	case "A Z":
	//		fmt.Println("Rock beats Scissors.")
	//		total += 3
	//	case "B Y":
	//		fmt.Println("Paper draws with Paper.")
	//		total += 5
	//	case "B X":
	//		fmt.Println("Paper beats Rock.")
	//		total += 1
	//	case "B Z":
	//		fmt.Println("Paper loses to Scissors.")
	//		total += 9
	//	case "C Y":
	//		fmt.Println("Scissors beats Paper.")
	//		total += 2
	//	case "C X":
	//		fmt.Println("Scissors lose to Rock.")
	//		total += 7
	//	case "C Z":
	//		fmt.Println("Scissors draw with Scissors.")
	//		total += 6
	//	default:
	//		fmt.Println("Not supposed to happen.")
	//	}
	//}

	for scanner.Scan() {
		strVal := scanner.Text()
		fmt.Println(strVal)
		switch outcome := strVal; outcome {
		case "A Y":
			fmt.Println("Rock - Draw.")
			total += 4
		case "A X":
			fmt.Println("Rock - Loss.")
			total += 3
		case "A Z":
			fmt.Println("Rock - Win.")
			total += 8
		case "B Y":
			fmt.Println("Paper - Draw.")
			total += 5
		case "B X":
			fmt.Println("Paper - Loss.")
			total += 1
		case "B Z":
			fmt.Println("Paper - Win.")
			total += 9
		case "C Y":
			fmt.Println("Scissors - Draw.")
			total += 6
		case "C X":
			fmt.Println("Scissors - Loss.")
			total += 2
		case "C Z":
			fmt.Println("Scissors - Win.")
			total += 7
		default:
			fmt.Println("Not supposed to happen.")
		}
	}

	fmt.Println("Total: ", total)
}
