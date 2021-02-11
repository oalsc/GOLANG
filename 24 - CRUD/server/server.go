package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CreateUser create a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read request body."))
		return
	}

	var user user

	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Error converting user to struct."))
		return
	}

	db, err := db.Conection()
	if err != nil {
		w.Write([]byte("Error connecting to db."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Error creating statement."))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(user.Nome, user.Email)
	if err != nil {
		w.Write([]byte("error when executing statement."))
		return
	}

	lastID, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("error getting the inserted id."))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted with id %d", lastID)))
}

// GetUsers get all users in database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conection()
	if err != nil {
		w.Write([]byte("Error to conect to database"))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Error in query db"))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user

		if err := rows.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.Write([]byte("Error in scan user"))
			return
		}

		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error to convert users to json"))
		return
	}
}

// GetUser get a user in database
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to convert string to uint"))
		return
	}

	db, err := db.Conection()
	if err != nil {
		w.Write([]byte("Error to conect to database"))
		return
	}

	row, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Error to select user with id"))
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.Write([]byte("Error to scan user"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error to encode user"))
		return
	}
}

// UpdateUser will update the user record in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to convert parameter to int"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error to read body of request"))
		return
	}

	var user user
	if erro := json.Unmarshal(body, &user); erro != nil {
		w.Write([]byte("Error to conver user to struct"))
		return
	}

	db, err := db.Conection()
	if err != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Error to creat statement"))
		return
	}

	if _, err := statement.Exec(user.Nome, user.Email, ID); err != nil {
		w.Write([]byte("Error to update user in database"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser will update user record in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to convert parameter to int"))
		return
	}

	db, err := db.Conection()
	if err != nil {
		w.Write([]byte("Error to connect to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Error to creat statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Error to Delete user in database"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
