package main

import (
	"net/http"
	"numbergame/backend/utils"
	"numbergame/backend/endpoints"
	"cloud.google.com/go/firestore"
)

var client *firestore.Client

func init() {
	utils.FirestoreInit()
}

func main() {
	http.HandleFunc("/getNumbers", utils.Logging(endpoints.GetNumbers))
	http.HandleFunc("/createUser", utils.Logging(endpoints.CreateUser))
	http.ListenAndServe(":80", nil)
}
