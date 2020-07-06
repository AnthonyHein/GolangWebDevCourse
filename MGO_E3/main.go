package main

import (
	"GolangWebDevCourse/MGO_E3/controllers"
	"net/http"
)



const sessionLength int = 30

func main() {
	c := controllers.NewController()
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.Signup)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
