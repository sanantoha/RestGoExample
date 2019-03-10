package main

import (
	"encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type UserEndpoint interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

type UserEndpointImpl struct {
	UserRepository UserRepository
}

func (ue UserEndpointImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Params:", params)

	var user User
	var err error
	if name, ok := params["name"]; ok {
		user, err = ue.UserRepository.GetUser(name)
		log.Println(name)
		if err != nil {
			log.Println(err)
		}		
	}	
	json.NewEncoder(w).Encode(user)	
}