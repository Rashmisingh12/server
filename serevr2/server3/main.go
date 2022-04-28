package main

import (
	"fmt"
	"net/http"
)

type people int

func (p people) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello  Rashmi Singh this side")
}

func main(){
	var p1 people
	http.ListenAndServe(":8080",p1)
}