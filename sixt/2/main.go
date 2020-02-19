package main

import "fmt"

func Solution_(blocks []int) int {
	longestDistance, currentDistance, lastLowestIndex, lastHighestIndex := 0, 0, 0, 0
	for i := 1; i < len(blocks); i++ {
		if blocks[i] >= blocks[i-1] {
			lastIndex := lastLowestIndex
			if lastLowestIndex > 0 {
				lastIndex = lastHighestIndex
			}

			currentDistance = i - lastIndex + 1
		}

		if blocks[i] > blocks[i-1] {
			lastHighestIndex = i
		}

		if blocks[i] < blocks[i-1] {
			currentDistance = i - lastHighestIndex + 1
			lastLowestIndex = i
		}

		if currentDistance > longestDistance {
			longestDistance = currentDistance
		}
	}

	return longestDistance
}

func Solution(blocks []int) int {
	longestDistance := 0
	currentDistance := 0
	lastDirection := 1
	direction := 1
	firstHalf := 1
	secondHalf := 0
	for i := 1; i < len(blocks); i++ {
		if (i > 1) && (blocks[i] < blocks[i-1] && blocks[i-1] > blocks[i-2]) {
			direction *= -1
			firstHalf--
			if firstHalf > 0 && secondHalf > 0 {
				currentDistance = firstHalf + secondHalf
				firstHalf, secondHalf = 0, 0

				if currentDistance > longestDistance {
					longestDistance = currentDistance
				}
			}
		}

		if (i > 1) && (blocks[i] > blocks[i-1] && blocks[i-1] < blocks[i-2]) {
			direction *= -1
			firstHalf--
		}

		if direction == lastDirection {
			firstHalf++
		}

		if direction == -1*lastDirection {
			secondHalf++
		}

		lastDirection = direction
	}

	currentDistance = firstHalf + secondHalf
	if currentDistance > longestDistance {
		return currentDistance
	}

	return longestDistance
}

func main() {
	r := Solution([]int{1, 5, 5, 2, 6})
	fmt.Println(r, r == 4)
	r = Solution([]int{1, 5, 5, 2, 6, 7, 7, 8})
	fmt.Println(r, r == 7)
	r = Solution([]int{2, 6, 8, 5})
	fmt.Println(r, r == 3)
	r = Solution([]int{1, 1})
	fmt.Println(r, r == 2)
	r = Solution([]int{1, 1, 2, 2})
	fmt.Println(r, r == 4)
	r = Solution([]int{1, 1, 3, 4, 3, 2, 4, 5, 6, 3})
	fmt.Println(r, r == 6)
	r = Solution([]int{1, 1, 3, 4, 3, 2, 4, 5, 6, 3, 2, 1, 1, 1, 1, 1, 2, 3, 4})
	fmt.Println(r, r == 1)

}
