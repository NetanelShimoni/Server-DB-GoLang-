package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)


type User struct {

	userid int `json:"userid"`
	name string `json:"name"`
	lastname string `json:"lastname"`
	age int `json:"age"`
	birthdate string `json:"birthdate"`


}


func getUsers(w http.ResponseWriter , r *http.Request){
	fmt.Sprint("asdsadsa")
	db := OpenConnection()
	rows,err := db.Query("SELECT * FROM temp_table")
	if err !=nil {
		log.Fatal(err)
	}
	 var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.userid, &user.name,&user.lastname,&user.age,&user.birthdate)
		users = append(users, user)
	}
	usersButes,_:= json.MarshalIndent(users,"","\t")
	w.Header().Set("Content-Type","application/json")
	w.Write(usersButes)
	defer rows.Close()
	defer db.Close()


}
func postUsers(writer http.ResponseWriter, request *http.Request) {
	db:= OpenConnection()
	fmt.Sprint("i am here")
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err !=nil {
		http.Error(writer,err.Error(), http.StatusBadRequest)
	}
	sqlStatment := `INSERT INTO temp_table (fname,lname) VALUES ($1,$2)`
	_,err = db.Exec(sqlStatment,user.userid,user.name,user.lastname,user.age,user.birthdate)
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
	sqlStatment := `UPDATE temp_table SET userid = $1, name=$2, lastname=$3, age=$4,birthdate=$5 WHERE userid=$6; `
	_,err = db.Exec(sqlStatment,user.userid,user.name,user.lastname,user.age,user.birthdate,vars["id"])
	if err !=nil {
		writer.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
	defer db.Close()
}


//Delete
func deleteUser(writer http.ResponseWriter, request *http.Request) {
	db:= OpenConnection()
	vars := mux.Vars(request)
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err !=nil {
		http.Error(writer,err.Error(), http.StatusBadRequest)
	}
	sqlStatment := `DELETE FROM temp_table WHERE id = $1;; `
	_,err = db.Exec(sqlStatment,vars["id"])
	if err !=nil {
		writer.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
	defer db.Close()
}


func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=122236 dbname=temp sslmode=disable")
	if err != nil {
		panic(err)
	}
	// check db
	err = db.Ping()

	fmt.Println("Connected!")
	return db
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", postUsers)
	http.HandleFunc("/users/{id}", updateUser)
	http.HandleFunc("/users/{id}", deleteUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}