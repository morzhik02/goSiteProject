package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"goSiteProject/app/model"
	"net/http"
)

func GetUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	//get list of all users
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	//return the list to the client in JSON format
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
}