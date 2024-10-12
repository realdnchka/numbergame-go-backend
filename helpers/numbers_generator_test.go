package helpers

import(
	"testing"
	"math/rand/v2"
)

// func TestNumberGenerator(t *testing.T) {
	
// }

func TestGenerateNumbersSumValidation(t *testing.T) {
	// Test case: totalNumbers = 10
	totalNumbers := 10
	numbers, expectedSum := GenerateNumbers(totalNumbers)

	// Check if the sum of a subset of numbers equals the expected sum
	calculatedSum := 0
	for _, num := range numbers {
		// Simulate the condition used in the original function (randomly add to sum)
		if rand.IntN(2) == 1 {
			calculatedSum += num
		}
	}

	// Edge case check: Ensure at least one non-zero sum
	if calculatedSum == 0 && totalNumbers > 1 {
		calculatedSum = numbers[rand.IntN(totalNumbers)]
	}

	// Compare the calculated sum with the expected sum
	if calculatedSum != expectedSum {
		t.Errorf("Expected sum %d, but got %d from the subset of numbers", expectedSum, calculatedSum)
	}
}