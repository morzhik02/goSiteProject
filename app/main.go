package main

import (
	"github.com/julienschmidt/httprouter"
	"goSiteProject/app/controller"

	//"goSiteProject/app/controller"
	"log"
	"net/http"
)

func main() {
	//create and launch a router to service requests
	router := httprouter.New()
	routes(router)

	//attach to the host and a free port for receiving and servicing incoming requests
	//the second parameter is passed to the router, which will work with requests
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(router *httprouter.Router) {
	//the path to the folder with external files: html, js, css, images, etc.
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	//what should be done for incoming requests of the specified type and at the specified address
	router.GET("/", controller.StartPage)
	router.GET("/users", controller.GetUsers)
}
