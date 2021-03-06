package adventofcode2015

import (
	"fmt"
	"strings"
)

func DayOne() {
	var input = ReadFile("day1-input.txt")

	fmt.Println(input)

	fmt.Println(len(input))

	ups := strings.Count(input, "(")
	downs := strings.Count(input, ")")
	total := ups - downs

	fmt.Printf("Ups: %v", ups)
	fmt.Printf("Downs: %v", downs)
	fmt.Printf("Final: %v", total)

	currentFloor := 0
	pos := 0

	for i := 0; i < len(input); i++ {

		if input[i:i+1] == "(" {
			currentFloor++
		}

		if input[i:i+1] == ")" {
			currentFloor--
		}
		if currentFloor == -1 {
			pos = i
			break
		}
	}

	fmt.Printf("Poss: %v", pos)

}
