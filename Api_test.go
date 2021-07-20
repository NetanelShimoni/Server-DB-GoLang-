package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test200ResponseGet(t *testing.T)  {
	request,_ := http.NewRequest("GET","/users",nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal (t,200,response.Code,"Test pass")
}
func Test200ResponsePost(t *testing.T)  {
	user := map[string]interface{}{
		"Userid": 255,
		"Name": "Alice",
		"Lastname": "Bob",
		"Age": 12,
		"Birthdate": "30-05-2021",
	}
	body, _ := json.Marshal(user)

	request,_ := http.NewRequest("POST","/users",bytes.NewReader(body))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal (t,200,response.Code,"Test pass")
}
func Test200ResponseUpdate(t *testing.T)  {
	user := map[string]interface{}{
		"Name": "david",
		"Lastname": "Bob",
		"Age": 12,
		"Birthdate": "30-05-2021",
	}
	body, _ := json.Marshal(user)

	request,_ := http.NewRequest("PUT","/users/255",bytes.NewReader(body))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal (t,200,response.Code,"Test pass")
}

//func Test200PostAndGet(t *testing.T)  {
////POST
//	user := map[string]interface{}{
//		"Userid": 255,
//		"Name": "Alice",
//		"Lastname": "Bob",
//		"Age": 12,
//		"Birthdate": "30-05-2021",
//	}
//	body, _ := json.Marshal(user)
//	request,_ := http.NewRequest("POST","/users",bytes.NewReader(body))
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response,request)
////GET
//
//	requestg,_ := http.NewRequest("GET","/users",nil)
//	responseg := httptest.NewRecorder()
//	Router().ServeHTTP(responseg,requestg)
//
//	//var data map[string]interface{}
//
//		opp, _ := json.Marshal(responseg.Body)
//
//	//fmt.Println(responseg.Body)
//
//	jsonMap := make(map[string]interface{})
//	err := json.Unmarshal([]byte(opp), &jsonMap)
//		if err !=nil {
//			panic(err)
//		}
//
//
//		fmt.Println(jsonMap)
//	//data := make(map[string]interface{})
//
//
//
//
//
//
//
//
//
//
//
//
//
////	db := OpenConnection()
////	rows,err := db.Query("SELECT * FROM user_table")
////	if err !=nil {
////		panic(err)
////	}
////
////	var users [] User
////
////	for rows.Next() {
////		var user User
////		rows.Scan(&user.Userid, &user.Name,&user.Lastname,&user.Age,&user.Birthdate)
////		if(user.Userid==255){
////			users = append(users, user)
////		}
////		//fmt.Println(user.birthdate)
////	}
////
////	json.Marshal(users)
////	 data := make(map[string]interface{})
////
////	fmt.Println(users)
////
////	//assert.Equal(t, user,users)
//}



