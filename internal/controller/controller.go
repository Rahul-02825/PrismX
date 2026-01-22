package controller

import (
	// "io"
	// "log"
	"net/http"
	"PrismX/internal/models"
	"PrismX/internal/database"
	"PrismX/logger"
	"encoding/json"

) 

func CreateUser(res http.ResponseWriter,req *http.Request){

	// Make sure it is a POST request
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}


	res.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil{
		logger.Instance.Error("Error in decoding json from request(user controller)")
		http.Error(res,err.Error(),http.StatusBadRequest)
		return 
	}

	result,err := database.UserCollection.InsertOne(req.Context(),user)
	
	if err != nil{
		logger.Instance.Error("Server error in creating user")
		http.Error(res,err.Error(),500)
		return
	}
	logger.Instance.Info("New user created successfully")
	json.NewEncoder(res).Encode(result)

}