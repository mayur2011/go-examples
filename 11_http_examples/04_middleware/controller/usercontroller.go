package controller

import (
	"encoding/json"
	"go-examples/11_http_examples/04_middleware/domain"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	//Dependencies & States
	Store domain.UserStore
}

// Get user will get all users
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	users, err := uc.Store.GetAll()
	if err != nil {
		return domain.User{}, http.StatusNotFound, err
	}
	return users, http.StatusOK, nil
}

// Post user will create a user
func (uc UserController) PostUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return domain.User{}, http.StatusBadRequest, err
	}
	if err := uc.Store.Create(user); err != nil {
		return domain.User{}, http.StatusNotFound, err
	}
	return user, 0, nil
}

// Delete user will delete given user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := uc.Store.Delete(id); err != nil {
		return domain.User{}, http.StatusNotFound, err
	}
	return nil, http.StatusAccepted, nil
}
