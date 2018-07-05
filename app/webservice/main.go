package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufemon/project3/app/users"
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
	user.Data.Id, _ = strconv.Atoi(r.FormValue("id"))
	user.Data.Name = r.FormValue("name")
	user.Data.Balance, _ = strconv.Atoi(r.FormValue("balance"))
	err = user.Insert(user.Data.Id, user.Data.Name, user.Data.Balance)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("{status:'ok',response:'insert success'}")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := users.New()
	if err != nil {
		log.Fatal(err)
	}
	user.Data.Id, _ = strconv.Atoi(r.FormValue("id"))
	user.Data.Name = r.FormValue("name")
	user.Data.Balance, _ = strconv.Atoi(r.FormValue("balance"))
	err = user.Update(user.Data.Id, user.Data.Name, user.Data.Balance)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("{status:'ok',response:'update success'}")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := users.New()
	if err != nil {
		log.Fatal(err)
	}
	user.Data.Id, _ = strconv.Atoi(r.FormValue("id"))
	err = user.Delete(user.Data.Id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("{status:'ok',response:'delete success'}")
}
