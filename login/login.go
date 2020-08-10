package login

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/ignite2405/login-regis/db"
	"github.com/ignite2405/login-regis/model"
)

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	var userLogin model.UserDetail
	var arrRes []model.UserDetail
	var response model.Response
	var username string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(reqBody, &userLogin)

	email := userLogin.Email
	shapass := sha256.Sum256([]byte(userLogin.Password))
	password := hex.EncodeToString(shapass[:])

	if email == "" || password == "" {
		response.Status = 0
		response.Message = "Provide the email and password for login"
	} else {

		dbcon := db.Dbconnect()
		users, err := dbcon.Query("select name from users where email = ? and password = ? limit 1", email, password)
		if err != nil {
			log.Fatal(err.Error())
		}

		for users.Next() {
			if err := users.Scan(&username); err != nil {
				log.Fatal(err.Error())
			}
		}

		if username == "" {
			response.Status = 0
			response.Message = "wrong email or password"
		} else {
			resDetail := model.UserDetail{Name: username, Email: email}

			arrRes = append(arrRes, resDetail)

			response.Status = 1
			response.Message = "Successfully Logged in"
			response.Data = arrRes
		}

		defer dbcon.Close()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
