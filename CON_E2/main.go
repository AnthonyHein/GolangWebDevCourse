package main

import (
    "net/http"
    "html/template"
    "log"
    "io"
)

func foo(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
    tpl := template.Must(template.ParseFiles("dog.gohtml"))
    err := tpl.Execute(res, nil)
    if err != nil {
        log.Fatalln(err)
    }
}

func doggo(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "dog.jpg")
}

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/dog/", dog)
    http.Handle("/dog.jpg", http.FileServer(http.Dir(".")))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
