package controller

import (
	"go-examples/11_http_examples/04_middleware/domain"
	"net/http"
)

type UserController struct {
	//Dependencies & States
	Store domain.UserStore
}

// Get user will get all users
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	return domain.User{}, 0, nil
}

// Post user will create a user
func (uc UserController) PostUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	return domain.User{}, 0, nil
}

// Delete user will delete given user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	return domain.User{}, 0, nil
}
