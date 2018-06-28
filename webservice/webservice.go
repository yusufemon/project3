package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yusufemon/project3/users"
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
	data := users.Get()
	json.NewEncoder(w).Encode(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var data users.Data
	json.NewDecoder(r.Body).Decode(&data)
	response := users.Insert(data.Id, data.Name, data.Balance)
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var data users.Data
	json.NewDecoder(r.Body).Decode(&data)
	response := users.Update(data.Id, data.Name, data.Balance)
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var data users.Data
	json.NewDecoder(r.Body).Decode(&data)
	response := users.Delete(data.Id)
	json.NewEncoder(w).Encode(response)
}