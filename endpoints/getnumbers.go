package endpoints

import (
	"encoding/json"
	"net/http"
	"numbergame/backend/utils"
	"numbergame/backend/utils/assert"
	"strconv"
)

//Endpoint for generating set of numbers and summary
func GetNumbers(w http.ResponseWriter, r *http.Request) {
	var countParam = r.URL.Query().Get("count")
	assert.NotEmptyString(countParam)
	
	var count int64 = 5
	
	count, err := strconv.ParseInt(countParam, 10, 32)
	assert.NotError(err, "Unable to parse 'count' param")
	
	var numbers, sum = utils.GenerateNumbers(int(count))
	
	numberResponse := utils.NumberResponse{
		Numbers: numbers,
		Sum: sum,
	}
	
	json.NewEncoder(w).Encode(numberResponse)
}