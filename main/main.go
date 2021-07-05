package main

import (
	"fmt"
	"net/http"
)

func main() {
	request1()
}

func homePage(reponse http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to the Go Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Claude Roy"
	fmt.Fprintf(response,"A little bit about me....")
	fmt.Println("Endpoint Hit: ",who)
}

func request1() {
	http.HandleFunc("/",homePage)
	http.HandleFunc("/aboutMe",aboutMe)
	http.ListenAndServe(":4000",nil)
}
