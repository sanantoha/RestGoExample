package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "errors"
    "net/http"
)

type UserEndpoint interface {
    GetUser(w http.ResponseWriter, r *http.Request)
    GetUsers(w http.ResponseWriter, r *http.Request)
    CreateUser(w http.ResponseWriter, r *http.Request)
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
            http.Error(w, errors.New("GetUser return error").Error(), 500)
        }       
    }   
    json.NewEncoder(w).Encode(user) 
}

func (ue UserEndpointImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := ue.UserRepository.GetUsers()
    if err != nil {
        log.Println(err)
        http.Error(w, errors.New("GetUsers returns error").Error(), 500)
    }
    json.NewEncoder(w).Encode(users)
}

func (ue UserEndpointImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    _ = json.NewDecoder(r.Body).Decode(&user)

    err := ue.UserRepository.InsertUser(&user)
    if err != nil {
        log.Println(err)
        http.Error(w, errors.New("CreateUser returns error").Error(), 500)
    }
    json.NewEncoder(w).Encode(user)
}