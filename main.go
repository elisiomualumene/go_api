package main

import "net/http"
import "fmt"

type User struct {
	id String
	name String
	created_at String
}

func main(){
	user := User{
		id: "1234",
		name: "Elisio Mualumene",
		created_at: "11/03/2023",
	}
	println(user)
	http.HandleFunc("/", HomeHandle)
	http.ListenAndServe(":3000",nil)
}

func HomeHandle(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}