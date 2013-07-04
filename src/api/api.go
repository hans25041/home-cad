package api

import (
  "fmt"
  "encoding/json"

  "appengine"
  "appengine/datastore"

  "net/http"
  "github.com/gorilla/mux"
  "strconv"
)

func init() {
  r := mux.NewRouter()
  r.HandleFunc("/api/{resource}",       create).Methods("POST")
  r.HandleFunc("/api/{resource}/{id}", get).Methods("GET")
  r.HandleFunc("/api/{resource}/{id}", update).Methods("PUT")
  r.HandleFunc("/api/{resource}/{id}", delete).Methods("DELETE")
  http.Handle("/", r)
}

type Room struct {
  Id   int
  Name string
}

func get_room(r *http.Request) (room Room, err error) {
  data := make([]byte, r.ContentLength)
  _,err = r.Body.Read(data)
  if err != nil {
    return
  }
  err = json.Unmarshal(data, &room)
  return
}

func get_vars(r *http.Request) (resource string, id int64, err error) {
  vars := mux.Vars(r)
  resource = vars["resource"]
  if id_str, has_id := vars["id"]; has_id {
    id, err := strconv.ParseInt(id_str, 10, 64)
  }
  return
}

func create(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  resource, _, err := get_vars(r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  room, err := get_room(r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  k := datastore.NewKey(c, "room", "0", room.Id)
  fmt.Fprintf(w,"%v",k)
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "room", nil), &room)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "Created %v:\n%v\nKey:%v\n", resource, room, key)
}

func get(w http.ResponseWriter, r *http.Request) {

  resource, id, err := get_vars(r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w,"Get %v %v\n", resource, id)

  c := appengine.NewContext(r)
  q := datastore.NewQuery("room").
    Filter("Id =", id)


  for t := q.Run(c); ; {
    var room Room
    key, err := t.Next(&room)
    if err == datastore.Done {
      break
    } else if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    fmt.Fprintf(w, "%v: %v", key, room)
  }

}

func update(w http.ResponseWriter, r *http.Request) {
  resource, id, err := get_vars(r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Fprintf(w,"Update %v %v", resource, id)
}

func delete(w http.ResponseWriter, r *http.Request) {
  resource, id, err := get_vars(r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Fprintf(w,"Delete %v %v", resource, id)
}
