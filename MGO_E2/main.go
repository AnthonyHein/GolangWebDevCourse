package main

import (
	"GolangWebDevCourse/MGO_E2/controllers"
	"GolangWebDevCourse/MGO_E2/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()

	// Load data.
	var users []models.User
	fmt.Println("Loading data.")

	bs, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatal(err)
    }

	err = json.Unmarshal(bs, &users)
	if err != nil {
        log.Fatal(err)
    }

	for _, elem := range users {
		uc.LoadUser(elem)
	}

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
