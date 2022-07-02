package main

import (
	"fmt"
	"github.com/bihari123/go-course/pkg/handlers"
	"net/http"

)

const portNumber = ":8080"

func main() {
	//fmt.Println("hello,World !")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil) //even if there is an error, i don't care. so put a _ here
}
