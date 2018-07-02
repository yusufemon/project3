package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/yusufemon/project3/users"

	"encoding/json"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users", UpdateUser).Methods("PUT")
	router.HandleFunc("/users", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user, err := users.New()
	if err != nil {
		log.Fatal(err)
	}
	data, err := user.Get()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := users.New()
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(r.Body).Decode(user.Data)
	response, err := user.Insert(user.Data.Id, user.Data.Name, user.Data.Balance)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := users.New()
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(r.Body).Decode(user.Data)
	response, err := user.Update(user.Data.Id, user.Data.Name, user.Data.Balance)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := users.New()
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(r.Body).Decode(user.Data)
	response, err := user.Delete(user.Data.Id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(response)
}
