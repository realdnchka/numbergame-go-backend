package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"numbergame/backend/utils"
)

type NumberResponse struct {
	Numbers []int `json:"numbers"`
	Sum int `json:"sum"`
} 

func getNumbers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	var numbers, sum = utils.GenerateNumbers(5)
	
	numberResponse := NumberResponse{
		Numbers: numbers,
		Sum: sum,
	}
	
	json.NewEncoder(w).Encode(numberResponse)
}

func main() {
	http.HandleFunc("/getNumbers", getNumbers)

	http.ListenAndServe(":80", nil)
}
