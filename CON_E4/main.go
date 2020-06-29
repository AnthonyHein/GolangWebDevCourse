package main

import (
    "net/http"
    "html/template"
    "log"
)

func main() {
    fs := http.FileServer(http.Dir("resources"))
    http.Handle("/resources/", http.StripPrefix("/resources", fs))
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
