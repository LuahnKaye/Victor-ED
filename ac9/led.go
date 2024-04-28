package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var value string
		fmt.Scanln(&value)

		numLEDs := calculateLEDs(value)
		fmt.Printf("%d leds\n", numLEDs)
	}
}

func calculateLEDs(value string) int {
	leds := []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

	totalLeds := 0
	for _, digit := range value {
		digitValue := int(digit - '0')
		totalLeds += leds[digitValue]
	}

	return totalLeds
}