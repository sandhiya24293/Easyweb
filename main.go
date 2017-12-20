package main

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"

)
type Name struct{
	Uname string
}

func GetUsersDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside user")
	var sandy Name

 
	err := json.NewDecoder(r.Body).Decode(&sandy)
	fmt.Println("response",sandy,err)
	err =json.NewEncoder(w).Encode(sandy)

	
}

func main() {
	var r = mux.NewRouter()
	fs := http.FileServer(http.Dir("./web/"))
	//http.Handle("/", fs)
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/",fs))




	r.HandleFunc("/postdata",GetUsersDetails)

	server := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
	
}

