package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Cred struct {
	Usertype string
	UserID   string
}

type Approve struct {
	Usertype    string
	UserID      string
	Transaction string
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func recollect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds Cred
	body, err := ioutil.ReadAll(r.Body)
	Handle(err)
	_ = json.Unmarshal(body, &creds)
	fmt.Println("Data is:", creds.UserID)
	json.NewEncoder(w).Encode(creds)
}

func approve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds Cred
	body, err := ioutil.ReadAll(r.Body)
	Handle(err)
	_ = json.Unmarshal(body, &creds)
	fmt.Println("Data is:", creds.UserID)
	json.NewEncoder(w).Encode(creds)
}

func main() {

	Port := "3000"
	http.HandleFunc("/recollect", recollect)
	http.HandleFunc("/approve", approve)
	fmt.Printf("Running on Port %s \n", Port)
	http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
}
