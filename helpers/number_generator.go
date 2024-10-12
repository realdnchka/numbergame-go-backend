package helpers

import (
	"math/rand/v2"
)

func GenerateNumbers(totalNumbers int) ([]int, int) {
	var numbers []int
	var sum int = 0
	
	for i:= 0; i < totalNumbers; i++ {
		var num = rand.IntN(51 - 1) + 1
		numbers = append(numbers, num)
		if rand.IntN(2) == 1 {
			sum += num
		}
		
		if sum == 0 {
			sum += numbers[rand.IntN(totalNumbers - 1) + 1]
		}
	}
	
	return numbers, sum
}