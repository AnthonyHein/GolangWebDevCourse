package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func main() {
    http.HandleFunc("/", count)
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.ListenAndServe(":8080", nil)
}

func count(w http.ResponseWriter, req *http.Request) {
    c, err := req.Cookie("my-cookie")
    if c == nil {
        http.SetCookie(w, &http.Cookie{
            Name: "my-cookie",
            Value: "1",
            Path: "/",
        })
        fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
        fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
        return
    }
    if err != nil {
    	http.Error(w, http.StatusText(400), http.StatusBadRequest)
    	return
    }
    oldVal, err := strconv.Atoi(c.Value)
    newVal := strconv.Itoa(oldVal + 1)
    http.SetCookie(w, &http.Cookie{
        Name: "my-cookie",
        Value: newVal,
        Path: "/",
    })
    fmt.Fprintln(w, "YOUR COOKIE:", c)
}

// Using cookies, track how many times a user has been to your website domain.
