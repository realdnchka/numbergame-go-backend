package endpoints

import (
	"encoding/json"
	"net/http"
	"numbergame/backend/utils"
	"strconv"
)

//Endpoint for generating set of numbers and summary
func GetNumbers(w http.ResponseWriter, r *http.Request) {
	var count, err = strconv.ParseInt(r.URL.Query().Get("count"), 0, 64)
	if err != nil {
		count = 5
	}
	
	var numbers, sum = utils.GenerateNumbers(int(count))
	
	numberResponse := utils.NumberResponse{
		Numbers: numbers,
		Sum: sum,
	}
	
	json.NewEncoder(w).Encode(numberResponse)
}