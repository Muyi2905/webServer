package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	fileServer:= http.FileServer(http.Dir("/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/form", formHandle)

	if err:=http.ListenAndServe(":8080", nil);err!= nil {
		log.Fatal(err)
	}
	
}