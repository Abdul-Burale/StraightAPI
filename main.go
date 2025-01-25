package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Player struct {
	Name  string `json:"Name"`
	Level int `json:"Level"`
}

var Players = []Player{
	{Name: "Thomas", Level: 60},
	{Name: "Abdi", 	 Level: 10},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the landing page")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/Hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GETAllPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EndPoint Hit: GETAllPlayers")
	json.NewEncoder(w).Encode(Players)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	fmt.Println("EndPoint Hit: Hello")
}

func main() {
	//	handleRequests()
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/GetPlayers", GETAllPlayers)
	http.ListenAndServe(":8080", nil)
}
