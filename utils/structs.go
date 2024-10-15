package utils

// Struct for numbers and summary of them for core game
type NumberResponse struct {
	Numbers []int `json:"numbers"`
	Sum int `json:"sum"`
} 