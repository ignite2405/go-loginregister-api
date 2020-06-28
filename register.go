package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func registerFunc(w http.ResponseWriter, r *http.Request) {
	var newUser UserDetail
	var arr_user []UserDetail
	var response Response
	var username string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(reqBody, &newUser)

	name := newUser.Name
	email := newUser.Email
	shapass := sha256.Sum256([]byte(newUser.Password))
	password := hex.EncodeToString(shapass[:])

	db := dbconnect()

	if name == "" || email == "" || password == "" {
		response.Status = 0
		response.Message = "Provide the name, email and password to register"
	} else {

		//check if email is already used
		users, err := db.Query("select name from users where email = ? limit 1", email)
		if err != nil {
			log.Fatal(err.Error())
		}

		for users.Next() {
			if err := users.Scan(&username); err != nil {
				log.Fatal(err.Error())
			}
		}

		if username != "" {
			response.Status = 0
			response.Message = "this email is already used"
		} else {
			regisQuery, err := db.Prepare("INSERT INTO users(name,email,password) VALUES(?,?,?)")
			if err != nil {
				log.Fatal(err.Error())
			}

			regisQuery.Exec(name, email, password)

			res := UserDetail{Name: name, Email: email}

			arr_user = append(arr_user, res)

			response.Status = 1
			response.Message = "Successfully Added New User"
			response.Data = arr_user
		}

		defer db.Close()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
