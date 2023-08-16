package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	longest := 0
	leftCursor := 0
	rightCursor := 0
	currentStringMap := map[string]int{}

	for rightCursor < len(s) {
		letterRight := fmt.Sprintf("%c", s[rightCursor])
		value, ok := currentStringMap[letterRight]
		currentStringMap[letterRight] = rightCursor
		if !ok || value < leftCursor {
			currentLength := rightCursor - leftCursor + 1

			if currentLength > longest {
				longest = currentLength
			}

			rightCursor++
			continue
		}
		leftCursor = value + 1
		rightCursor++
	}

	return longest
}

func lengthOfLongestSubstringBF(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	longest := 0
	cursor := 0
	currentLength := 0
	currentStringMap := map[string]int{}

	for cursor < len(s) {
		letter := fmt.Sprintf("%c", s[cursor])
		value, ok := currentStringMap[letter]
		if !ok {
			currentStringMap[letter] = cursor
			currentLength++

			if currentLength > longest {
				longest = currentLength
			}

			cursor++
			continue
		}

		cursor = value + 1
		currentStringMap = map[string]int{}
		currentLength = 0
	}
	return longest
}
