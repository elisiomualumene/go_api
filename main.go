package main

import "net/http"

type User struct {
	id String
	name String
	created_at String
}

func main(){
	http.HandleFunc("/", HomeHandle)
	http.ListenAndServe(":3000",nil)
}

func HomeHandle(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func UserHandle(w, http.ResponseWriter, r *http.Request){
	
}