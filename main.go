package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}
type JamesBondPhrase struct {
    Phrase  string   `json:"jamesbondphrase,omitempty"`
}

var people []Person
var phrases []JamesBondPhrase

func GetJamesBondPhrase(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    var phrase JamesBondPhrase
    for _, item := range people {
        if item.ID == params["id"] {
            JamesBond := item.Lastname + ", " + item.Firstname + " " + item.Lastname
            phrase.Phrase = JamesBond
            // phrases = append(phrases, phrase)
            json.NewEncoder(w).Encode(phrase)
            return
        }
    }
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}
// our main function
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
    people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
    people = append(people, Person{ID: "4", Firstname: "Anna", Lastname: "Heikkila", Address: &Address{City: "Solna", State: "Uppland"}})
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    router.HandleFunc("/people/jb/{id}", GetJamesBondPhrase).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))



}




