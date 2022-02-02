package main

import "fmt"
import "net/http"

import "github.com/gorilla/mux"
import "rsc.io/quote"


func main() {
    // printing on console
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    // create a mux router and register the home page func
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello")})

    // run server on 10001 port
    http.Handle("/", r)
    http.ListenAndServe(":10001", nil)
}
