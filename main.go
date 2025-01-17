package main 

import (
    "fmt"
    "net/http"
    "log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful\n")
    name := r.FormValue("name")
    email := r.FormValue("email")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Email = %s\n", email)
}

func main() {
    fileServer := http.FileServer(http.Dir("static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/form", formHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}