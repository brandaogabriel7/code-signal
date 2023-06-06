package main

import "fmt"

func main() {
	fmt.Println(solution("foo(bar(baz))blim"))
}

func solution(inputString string) string {
	output, _ := getStringPortion(inputString, 0, false)
	return string(output)
}

func getStringPortion(inputString string, startIndex int, reverse bool) ([]rune, int) {
	portion := []rune{}
	for index := startIndex; index < len(inputString); index++ {
		char := inputString[index]
		if char == ')' {
			return portion, index
		}
		newPortion := []rune{}
		if char == '(' {
			newPortion, index = getStringPortion(inputString, index+1, !reverse)
		} else {
			newPortion = []rune{rune(char)}
		}
		if reverse {
			portion = append(newPortion, portion...)
		} else {
			portion = append(portion, newPortion...)
		}
	}
	return portion, len(inputString)
}
