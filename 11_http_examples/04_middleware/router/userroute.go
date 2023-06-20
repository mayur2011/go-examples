package router

import (
	"go-examples/11_http_examples/04_middleware/controller"
	"go-examples/11_http_examples/04_middleware/mapstore"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	mapStore := mapstore.NewUserMapStore()
	userController := controller.UserController{Store: mapStore}
	router.Handle("/create", controller.ResponseHandler(userController.PostUser)).Methods("POST")
	router.Handle("/getAll", controller.ResponseHandler(userController.GetUsers)).Methods("GET")
	router.Handle("/delete/{id}", controller.ResponseHandler(userController.DeleteUser)).Methods("DELETE")
	return router
}

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetUserRoutes(router)
	return router
}
