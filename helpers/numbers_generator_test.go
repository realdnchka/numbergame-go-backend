package helpers

import(
	"testing"
	"strconv"
)

func TestGenerateNumberCorrectTotalOfNumbers(t *testing.T) {
	numbers, sum := GenerateNumbers(5)
	_ = sum
	
	if(len(numbers) != 5) {
		t.Errorf("Total of numbers is incorrect. Total is %d, needed is %d", len(numbers), 5)
	}
}

func TestGenerateNumberMoreThanOneSolution(t *testing.T) {
	numbers, sum := GenerateNumbers(5)
	
	for i:= 0; i < len(numbers); i++ {
		if numbers[i] == sum {
			t.Errorf("\nOne solution discover! Number is: %d, index is: %d, slice is: %ssummary is: %d", numbers[i], i, convertArrayIntToString(numbers), sum)
		}
	}
}

func TestGenerateNumbersSumValidation(t *testing.T) {
	numbers, sum := GenerateNumbers(5)
	
	if !canSumSubset(GenerateNumbers(5)) {
		t.Errorf("\nWanted %d, numbers are: %s", sum, convertArrayIntToString(numbers))
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