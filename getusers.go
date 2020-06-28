package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users UserDetail
	var arr_user []UserDetail
	var response Response

	db := dbconnect()
	defer db.Close()

	rows, err := db.Query("Select name,email from users") // ,password
	if err != nil {
		//log.Print(err)
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
