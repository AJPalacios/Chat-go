package main

import(
  	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"sync"
)

type Response struct{

	Message  string "json:'message'"
	Status bool "json:'status'"
	IsValid bool "json:'IsValid'"


}

var Users = struct{
	m map[string] User
	sync.RWMutex
}{m: make(map[string] User)}

type User struct{
	user_name string
}

func main() {

	cssHandle := http.FileServer(http.Dir("./front/css/"))
	jsHandle := http.FileServer(http.Dir("./front/js/"))

	mux := mux.NewRouter()
	mux.HandleFunc("/Hola", HolaMundo).Methods("GET")
	mux.HandleFunc("/HolaJson", HolaMundoJson).Methods("GET")
	mux.HandleFunc("/Static", HolaHtml).Methods("GET")
	mux.HandleFunc("/validate", Validate).Methods("POST")

	http.Handle("/", mux)
	http.Handle("/css/", http.StripPrefix("/css/", cssHandle))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandle))
	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte ("Hola mundo Web") )
}

func HolaMundoJson(w http.ResponseWriter, r *http.Request) {
	response := CreateResponse("Esto esta en formato JSon", true)
	json.NewEncoder(w).Encode(response)
}

func CreateResponse(message string, valid bool) Response {
	return Response{message, valid, true}
}

func HolaHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./front/index.html")
}

func Validate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user_name := r.FormValue("nombre")

	response := Response{}

	if user_exist(user_name) {
		response.IsValid = false
	}else{
		response.IsValid = true
	}

	json.NewEncoder(w).Encode(response)
}

func user_exist(user_name string)bool{
	Users.RLock()
	defer Users.RUnlock()

	if _, ok :=Users.m[user_name]; ok{
		return true
	}
	return false
		
	
}
