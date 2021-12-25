package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func StartPage(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	text := "Hello! Welcome to the start page"
	//return text
	fmt.Fprint(writer, text)
}
