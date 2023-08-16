package twosum

// https://leetcode.com/problems/two-sum/

func twoSum(input []int, target int) []int {
	hashMap := map[int]int{}

	for index, value := range input {
		compoIndex, ok := hashMap[value]
		if !ok {
			component := target - value
			hashMap[component] = index
			continue
		}
		return []int{compoIndex, index}
	}

	return nil
}
