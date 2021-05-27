package main

import (
	"fmt"
	"strconv"
)

func main() {
	lowerBound := 172930
	upperBound := 683082

	fmt.Println("Puzzle 1:", countValidPasswords(lowerBound, upperBound, isValidPassword))
	fmt.Println("Puzzle 2:", countValidPasswords(lowerBound, upperBound, isValidPassword2))
}

func countValidPasswords(lowerBound, upperBound int, criteria func(int) bool) int {
	var count int

	for i := lowerBound; i <= upperBound; i++ {
		if criteria(i) {
			count += 1
		}
	}

	return count
}

func isValidPassword(input int) bool {
	runes := []rune(strconv.Itoa(input))
	return isNonDecreasingSeries(runes) && containsAdjacentDigits(runes)
}

func isValidPassword2(input int) bool {
	runes := []rune(strconv.Itoa(input))
	return isNonDecreasingSeries(runes) && containsAdjacentDigitsStrict(runes)
}

func containsAdjacentDigitsStrict(runes []rune) bool {
	groupCounter := 1
	prev := runes[0]

	for i, curr := range runes[1:] {
		if curr == prev {
			groupCounter += 1

			// check adjCounter is 2 for the last rune
			if i+1 == len(runes)-1 && groupCounter == 2 {
				return true
			}
		} else {
			// check adjCounter is 2 when digit has changed
			if groupCounter == 2 {
				return true
			}
			groupCounter = 1
		}
		prev = curr
	}
	return false
}

func containsAdjacentDigits(runes []rune) bool {
	prev := runes[0]

	for _, curr := range runes[1:] {
		if prev == curr {
			return true
		}

		prev = curr
	}

	return false

}

func isNonDecreasingSeries(runes []rune) bool {
	prev := runes[0]

	for _, curr := range runes[1:] {
		if int(curr-'0') < int(prev-'0') {
			return false
		}

		prev = curr
	}

	return true
}
