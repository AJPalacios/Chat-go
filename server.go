package main

import(
  "github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Response struct{

	Message  string "json:'message'"

}

func main() {

	cssHandle := http.FileServer(http.Dir("./front/css/")) 

	mux := mux.NewRouter()
	mux.HandleFunc("/Hola", HolaMundo).Methods("GET")
	mux.HandleFunc("/HolaJson", HolaMundoJson).Methods("GET")
	mux.HandleFunc("/Static", HolaHtml).Methods("GET")

	http.Handle("/", mux)
	http.Handle("/css/", http.StripPrefix("/css/", cssHandle))
	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte ("Hola mundo Web") )
}

func HolaMundoJson(w http.ResponseWriter, r *http.Request) {
	response := CreateResponse()
	json.NewEncoder(w).Encode(response)
}

func CreateResponse() Response {
	return Response{"Esto esta en formato JSon"}
}

func HolaHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./front/index.html")
}