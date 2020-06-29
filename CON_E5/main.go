package main

import (
    "net/http"
    "html/template"
    "log"
)

func main() {
    http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
    http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
    tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
    err := tpl.Execute(res, nil)
    if err != nil {
        log.Fatalln(err)
    }
}
