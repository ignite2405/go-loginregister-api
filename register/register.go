package register

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

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	var newUser model.UserDetail
	var arr_user []model.UserDetail
	var response model.Response
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

	dbcon := db.Dbconnect()

	if name == "" || email == "" || password == "" {
		response.Status = 0
		response.Message = "Provide the name, email and password to register"
	} else {

		//check if email is already used
		users, err := dbcon.Query("select name from users where email = ? limit 1", email)
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
			regisQuery, err := dbcon.Prepare("INSERT INTO users(name,email,password) VALUES(?,?,?)")
			if err != nil {
				log.Fatal(err.Error())
			}

			regisQuery.Exec(name, email, password)

			res := model.UserDetail{Name: name, Email: email}

			arr_user = append(arr_user, res)

			response.Status = 1
			response.Message = "Successfully Added New User"
			response.Data = arr_user
		}

		defer dbcon.Close()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
