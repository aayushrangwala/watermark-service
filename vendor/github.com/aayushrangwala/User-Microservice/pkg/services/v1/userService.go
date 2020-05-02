package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aayushrangwala/User-Microservice/util"
)

// GetUserProfile returns the profile of the user
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("Username")
	userDet := &util.UserDetails{
		Name:  userName,
		Dob:   "01/01/2000",
		Age:   19,
		Email: "ex@ex.com",
		Phone: "9999999999",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userDet); err != nil {
		panic(err)
	}
}

// GetMicroserviceName is the unsecured API and will send the microservice name
func GetMicroserviceName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("User Microservice")
	if err != nil {
		log.Fatal("Error encoding or returning the response", err)
	}
}

// GetUserEmail returns the email of the user
func GetUserEmail(w http.ResponseWriter, r *http.Request) {
	userEmail := "ex@ex.com"
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(userEmail))
	if err != nil {
		log.Fatal("Error wrting or returning the response", err)
	}
}
