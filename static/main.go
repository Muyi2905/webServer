package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found",  http.StatusAccepted)
	}

	if r.Method != "GET"{
		http.Error(w, "Method is not supported", http.StatusAccepted)
	

fmt.Fprintf(w, "hello")
	}
}

func main(){
	fileServer:= http.FileServer(http.Dir("/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err:=http.ListenAndServe(":8080", nil);err!= nil {
		log.Fatal(err)
	}
	
}