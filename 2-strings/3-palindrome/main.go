package main

import (
	"strings"
	"unicode"
)

//https://leetcode.com/problems/valid-palindrome-ii/
func validPalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}

	isOriginalPalidrome, leftIndex, rightIndex := isPalindrome(s)
	if isOriginalPalidrome {
		return true
	}

	// remove leftIndex
	newStringRmLeft := s[leftIndex+1:rightIndex+1]
	if ok, _, _ := isPalindrome(newStringRmLeft); ok {
		return true
	}

	newStringRmRight := s[leftIndex:rightIndex]
	if ok, _, _ := isPalindrome(newStringRmRight); ok {
		return true
	}
	return false
}

//https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) (bool, int, int) {
	if len(s) <= 1 {
		return true, 0,0 
	}
	leftIndex := 0
	rightIndex := len(s) - 1

	for leftIndex < rightIndex {
		leftValue := string(s[leftIndex])
		rightValue := string(s[rightIndex])
		if !isAlphanumeric(leftValue) {
			leftIndex++
			continue
		}
		if !isAlphanumeric(rightValue) {
			rightIndex--
			continue
		}
		if strings.EqualFold(leftValue, rightValue) {
			leftIndex++
			rightIndex--
		} else {
			return false, leftIndex, rightIndex
		}
	}
	return true, 0,0 
}

func isAlphanumeric(s string) bool {
	return unicode.IsLetter(rune(s[0])) || unicode.IsDigit(rune(s[0]))
}
