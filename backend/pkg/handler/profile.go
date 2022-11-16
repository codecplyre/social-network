package handler

import (
	"encoding/json"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/response"
	"backend/pkg/structs"
)

// Profile handles all requests for a users own profile information.
func (DB *Env) Profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "GET" {
		c, err := r.Cookie("session_token") // Check if a cookie is present
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("No cookie present user unauthorized", "Unauthorised", w)
			return
		}
		userID := r.URL.Query().Get("userID") // Get the parameter
		var result structs.User
		getErr := auth.GetUser("userId", userID, &result, *DB.Env)
		if getErr != nil {
			response.WriteMessage("Error getting user: "+getErr.Error(), "User Not Found", w)
			return
		}
		// for i := 0; i < 5; i++ {
		// 	newFollower := structs.User{
		// 		FirstName:   "HEllo",
		// 		LastName:    "World",
		// 		NickName:    "Nick",
		// 		Email:       strings.ToLower("email@" + uuid.NewV4().String() + ".com"),
		// 		Password:    "Password123",
		// 		DateOfBirth: "06-08-2002",
		// 		Avatar:      "images/profile/default-user.svg",
		// 	}
		// 	auth.InsertUser(newFollower, *DB.Env)
		// 	var result structs.User
		// 	auth.GetUser("email", newFollower.Email, &result, *DB.Env)
		// 	follow.InsertFollow(result.UserId, userID, DB.Env)
		// 	if i < 2 {
		// 	closefriend.AddCloseFriend(userID, result.UserId, *DB.Env)
		// 	}

		// }

		result.Password = ""
		profileDetails, marshErr := json.Marshal(result)
		if marshErr != nil {
			response.WriteMessage("Error marshalling user profile data", "500 Internal Server Error", w)
			return
		}
		w.Write(profileDetails)
	}
}
