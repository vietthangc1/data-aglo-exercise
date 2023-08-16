package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	sPointer := len(s) - 1
	tPointer := len(t) - 1

	sBack := 0
	tBack := 0

	for sPointer >= 0 || tPointer >= 0 {
		sLetter := ""
		if sPointer >= 0 {
			sLetter = fmt.Sprintf("%c", s[sPointer])

			if sLetter == "#" {
				sBack++
				sPointer--
				continue
			}

			if sBack > 0 {
				sPointer--
				sBack--
				continue
			}
		}

		tLetter := ""
		if tPointer >= 0 {
			tLetter = fmt.Sprintf("%c", t[tPointer])
			if tLetter == "#" {
				tBack++
				tPointer--
				continue
			}

			if tBack > 0 {
				tPointer--
				tBack--
				continue
			}
		}

		if sLetter == tLetter {
			sPointer--
			tPointer--
			continue
		} else {
			return false
		}
	}
	return true
}

func backspaceCompareBF(s string, t string) bool {
	newS := convertString(s)
	newT := convertString(t)
	return newS == newT
}

func convertString(s string) string {
	newS := ""
	for _, value := range s {
		letter := fmt.Sprintf("%c", value)
		if letter != "#" {
			newS += letter
			continue
		}
		if len(newS) == 0 {
			continue
		}
		newS = newS[:len(newS)-1]
	}
	return newS
}
