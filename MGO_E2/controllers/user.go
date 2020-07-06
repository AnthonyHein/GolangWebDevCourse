package controllers

import (
	"GolangWebDevCourse/MGO_E2/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"io/ioutil"
)

type UserController struct {
	Session map[bson.ObjectId]models.User
}

func NewUserController() *UserController {
	return &UserController{
		Session : make(map[bson.ObjectId]models.User),
	}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oid := bson.ObjectIdHex(id)

	// Fetch user
	if val, ok := uc.Session[oid]; ok {
		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(val)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "%s\n", uj)
	} else {
		w.WriteHeader(404)
		return
	}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = bson.NewObjectId()

	// store the user in map
	uc.Session[u.Id] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

	bs, err := json.Marshal(&uc.Session)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("log.json", bs, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Fetch user
	if _, ok := uc.Session[oid]; !ok {
		w.WriteHeader(404)
		return
	}

	delete(uc.Session, oid)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")

	bs, err := json.Marshal(&uc.Session)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("log.json", bs, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func (uc UserController) LoadUser(u models.User) {

	uc.Session[u.Id] = u
	uj, _ := json.Marshal(u)
	fmt.Printf("Loaded user:\n%s\n", uj)
}
