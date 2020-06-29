package main

import (
    "net/http"
)

func a(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("Index."))
}

func b(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("Dog."))
}

func c(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("Anthony Hein"))
}


func main() {
    http.HandleFunc("/", a)
    http.HandleFunc("/dog/", b)
    http.HandleFunc("/me/", c)
    http.ListenAndServe(":8080", nil)
}
