package getusers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/ignite2405/login-regis/db"
	"github.com/ignite2405/login-regis/model"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users model.UserDetail
	var arr_user []model.UserDetail
	var response model.Response

	dbcon := db.Dbconnect()
	defer dbcon.Close()

	rows, err := dbcon.Query("Select name,email from users") // ,password
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&users.Name, &users.Email /*, &users.Password*/); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
