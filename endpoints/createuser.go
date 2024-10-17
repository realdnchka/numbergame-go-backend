package endpoints

import(
	"net/http"
	"numbergame/backend/utils"
	"encoding/json"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user utils.User
	
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	
	// if user.Name == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }
	
	utils.Client.Collection("users").Doc(user.Name).Set(ctx, map[string]interface{}{
		        "highscore": 0,
		        "totalscores":  0,
		})
}
