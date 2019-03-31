package main

import (
    "fmt"
    "log"
    "github.com/BurntSushi/toml"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "database/sql"
)


func main() {

    var config Config

    if _, err := toml.DecodeFile("config.toml", &config); err != nil {
        log.Panic("Can not read config.toml ", err)
    }
    log.Println("Read config:", config)

    dbConf := config.DatabaseConfig

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable", 
        dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Panic(err)
    }
    defer db.Close()

    var userRepository UserRepository = &UserRepositoryImpl{db}
    var userEndpoint UserEndpoint = &UserEndpointImpl{userRepository}
    
    router := mux.NewRouter()
    router.HandleFunc("/users/{name}", userEndpoint.GetUser).Methods("GET")
    router.HandleFunc("/{users:users(?:\\/)?}", userEndpoint.GetUsers).Methods("GET")
    router.HandleFunc("/{users:users(?:\\/)?}", userEndpoint.CreateUser).Methods("POST")
    router.HandleFunc("/{users:users(?:\\/)?}", userEndpoint.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{name}", userEndpoint.DeleteUser).Methods("DELETE")
    port := strconv.Itoa(config.ServerConfig.Port)
    log.Println("Try run server on port:", port)
    log.Fatal(http.ListenAndServe(":" + port, router))
}
