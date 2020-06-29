package main

import (
    "net/http"
    "html/template"
    "log"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func a(res http.ResponseWriter, req *http.Request) {
    data := struct{
        First string
        Last string
        FavNums []int
    }{
        First : "Anthony",
        Last : "Hein",
        FavNums : []int{57, 8, 64, 125},
    }

    err := tpl.Execute(res, data)
    if err != nil {
        log.Fatalln(err)
    }
}

func main() {
    http.HandleFunc("/", a)
    http.ListenAndServe(":8080", nil)
}
