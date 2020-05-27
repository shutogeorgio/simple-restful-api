package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Word struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

var Words []Word

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllWords(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllWords")
	json.NewEncoder(w).Encode(Words)
}

func returnSingleWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, word := range Words {
		if word.Id == key {
			json.NewEncoder(w).Encode(word)
		}
	}
}

func createNewWord(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var word Word
	json.Unmarshal(reqBody, &word)
	Words = append(Words, word)
	json.NewEncoder(w).Encode(word)
}

func deleteWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, word := range Words {
		if word.Id == id {
			Words = append(Words[:index], Words[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/words", returnAllWords)
	myRouter.HandleFunc("/word", createNewWord).Methods("POST")
	myRouter.HandleFunc("/word/{id}", deleteWord).Methods("DELETE")
	myRouter.HandleFunc("/word/{id}", returnSingleWord)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Words = []Word{
		Word{Id: "1", Name: "Hello"},
		Word{Id: "2", Name: "Hello"},
	}
	handleRequests()
}
