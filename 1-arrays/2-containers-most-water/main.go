package containersmostwater

// https://leetcode.com/problems/container-with-most-water/

// brute force: loop over all element pairs and find max area
func maxArea(height []int) int {
	max := 0

	currentArea := 0

	leftIndex := 0
	rightIndex := len(height) - 1

	for leftIndex < rightIndex {
		minHeight, side := min(height[leftIndex], height[rightIndex])
		currentArea = calculateArea(leftIndex, rightIndex, minHeight)
		if currentArea > max {
			max = currentArea
		}
		if side == "left" {
			leftIndex++
			continue
		}
		rightIndex--
	}

	return max

}

func min(a, b int) (int, string) {
	if a < b {
		return a, "left"
	}
	return b, "right"
}

func calculateArea(leftIndex, rightIndex, height int) int {
	return height * (rightIndex - leftIndex)
}
