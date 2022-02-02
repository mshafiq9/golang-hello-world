package main

import "fmt"

import "rsc.io/quote"
import "github.com/gorilla/mux"
import "net/http"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello")})

    http.Handle("/", r)
    http.ListenAndServe(":10001", nil)
}
