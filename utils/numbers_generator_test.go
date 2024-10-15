package utils

import(
	"testing"
	"strconv"
)

var arr = []int {3, 4, 5, 6, 7, 8, 9, 10}

func TestGenerateNumberCorrectTotalOfNumbers(t *testing.T) {
	for _, count := range arr {
		t.Run("Generates correct total of numbers", func(t *testing.T) {
			numbers, _ := GenerateNumbers(count)
			
			if(len(numbers) != count) {
				t.Errorf("Total of numbers is incorrect. Got: %d, expected: %d", len(numbers), count)
			}
		})
	}
}

func TestGenerateNumberMoreThanOneSolution(t *testing.T) {
	for _, count := range arr {
		t.Run("Not generates one-num solution", func(t *testing.T) {
			numbers, sum := GenerateNumbers(count)
			
			for i:= 0; i < len(numbers); i++ {
				if numbers[i] == sum {
					t.Errorf("\nOne-num solution discover! Number is: %d, index is: %d, slice is: %ssummary is: %d", numbers[i], i, convertArrayIntToString(numbers), sum)
				}
			}
		})
	}
}

func TestGenerateNumbersSumValidation(t *testing.T) {
	for _, count := range arr {
		t.Run("Sum validation", func(t *testing.T) {
			numbers, sum := GenerateNumbers(count)
			
			if !canSumSubset(GenerateNumbers(count)) {
				t.Errorf("\nWanted %d, numbers are: %s", sum, convertArrayIntToString(numbers))
			}
		})
	}
}

func convertArrayIntToString(arr []int) string {
	output := ""
	for _, v := range arr {
		temp := strconv.Itoa(v)
		output = output + temp + ", "
	}
	
	return output
}

// Helper function: Subset sum check
func canSumSubset(numbers []int, sum int) bool {
	// dp[j] will be true if a subset with sum j can be formed using the numbers
	dp := make([]bool, sum+1)
	dp[0] = true // Base case: sum of 0 is always achievable

	for _, num := range numbers {
		// Traverse backwards to avoid recomputing results in the same iteration
		for j := sum; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[sum] // Return true if we can form the target sum
}