package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)



type User struct {

	Userid int `json:"userid"`
	Name string `json:"name"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
	Birthdate string `json:"birthdate"`


}


func getUsers(w http.ResponseWriter , r *http.Request){
	db := OpenConnection()
	rows,err := db.Query("SELECT * FROM user_table")
	if err !=nil {
		panic(err)
	}

	var users [] User

	for rows.Next() {
		var user User
		rows.Scan(&user.Userid, &user.Name,&user.Lastname,&user.Age,&user.Birthdate)
		//fmt.Println(user.birthdate)
		users = append(users, user)
	}
	//fmt.Println(users)
	w.Header().Set("Content-type", "application/json")
	jData, err := json.MarshalIndent(users," ","\t")
	w.Write(jData)
	defer rows.Close()
	defer db.Close()

}
func postUsers(writer http.ResponseWriter, request *http.Request) {
	db:= OpenConnection()
	fmt.Sprint("i am here")
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err !=nil {
		http.Error(writer,err.Error(),  http.StatusBadRequest)
	}


	sqlStatment := `INSERT  INTO user_table (userid,name,lastname,age,birthdate) VALUES ($1,$2,$3,$4,$5)`
	_,err = db.Exec(sqlStatment,user.Userid,user.Name,user.Lastname,user.Age,user.Birthdate)
	if err !=nil {
		writer.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
	defer db.Close()
}

//UPDATE (PUT)
func updateUser(writer http.ResponseWriter, request *http.Request) {
	db:= OpenConnection()
	vars := mux.Vars(request)
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err !=nil {
		http.Error(writer,err.Error(), http.StatusBadRequest)
	}
	sqlStatment := `UPDATE user_table SET userid = $1, name=$2, lastname=$3, age=$4,birthdate=$5 WHERE userid=$6; `
	_,err = db.Exec(sqlStatment,vars["id"],user.Name,user.Lastname,user.Age,user.Birthdate,vars["id"])
	if err !=nil {
		writer.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
	defer db.Close()
}


//Delete
func deleteUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hiii!!!")
	db:= OpenConnection()
	vars := mux.Vars(request)
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err !=nil {
		http.Error(writer,err.Error(), http.StatusBadRequest)
	}
	sqlStatment := `DELETE FROM user_table WHERE userid = $1;; `
	_,err = db.Exec(sqlStatment,vars["id"])
	if err !=nil {
		writer.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
	defer db.Close()
}


func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=122236 dbname=users sslmode=disable")
	if err != nil {
		panic(err)
	}
	// check db
	err = db.Ping()
	fmt.Println("Connected!")
	return db
}
