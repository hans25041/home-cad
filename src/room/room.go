package room

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
)

func init() {
  r := mux.NewRouter()
  r.HandleFunc("/room",       create).Methods("POST")
  r.HandleFunc("/room/{key}", get).Methods("GET")
  r.HandleFunc("/room/{key}", update).Methods("PUT")
  r.HandleFunc("/room/{key}", delete).Methods("DELETE")
//  http.Handle("/", r)
}

func create(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"Create Room")
}

func get(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["key"]
  fmt.Fprintf(w,"Get Room %v", key)
}

func update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["key"]
  fmt.Fprintf(w,"Update Room %v", key)
}

func delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["key"]
  fmt.Fprintf(w,"Delete Room %v", key)
}
