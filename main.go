package main

import (
	"net/http"
	"numbergame/backend/utils"
	"numbergame/backend/endpoints"
)

func main() {
	http.HandleFunc("/getNumbers", utils.Logging(endpoints.GetNumbers))

	http.ListenAndServe(":80", nil)
}
