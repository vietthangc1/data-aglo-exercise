package main

import (
	"log"
	"math"
)

var (
	input  = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expect = 6
)

//https://leetcode.com/problems/trapping-rain-water/
func main() {
	out := trap(input)
	log.Println(out)
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	leftIndex := 0
	rightIndex := len(height) - 1

	leftMax := 0
	rightMax := 0

	totalWater := 0

	for leftIndex < rightIndex {
		leftValue := height[leftIndex]
		rightValue := height[rightIndex]

		if leftValue <= rightValue {
			if leftIndex == 0 {
				leftMax = leftValue
				leftIndex++
				continue
			}

			if leftMax < leftValue {
				leftMax = leftValue
			} else {
				currentWater := leftMax - leftValue
				totalWater += currentWater
			}
			leftIndex++
		} else {
			if rightIndex == 0 {
				rightMax = rightValue
				rightIndex--
				continue
			}

			if rightMax < rightValue {
				rightMax = rightValue
			} else {
				currentWater := rightMax - rightValue
				totalWater += currentWater
			}
			rightIndex--
		}
	}

	return totalWater
}

func trapBruteForce(height []int) int {
	lenHeight := len(height)
	if lenHeight < 3 {
		return 0
	}

	totalWater := 0
	rightMaxAtIndex := make([]int, lenHeight)

	currentMax := 0
	for i := lenHeight - 1; i >= 0; i-- {
		if height[i] > currentMax {
			currentMax = height[i]
		}
		rightMaxAtIndex[i] = currentMax
	}

	currentLeftMax := 0
	for i := 0; i < lenHeight; i++ {
		if height[i] > currentLeftMax {
			currentLeftMax = height[i]
		}

		totalWater += int(math.Min(float64(currentLeftMax), float64(rightMaxAtIndex[i])) - float64(height[i]))
	}
	return totalWater
}
