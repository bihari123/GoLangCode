package main

import (
	"fmt"
	"net/http"

	"github.com/tarun/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	//fmt.Println("hello,World !")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil) //even if there is an error, i don't care. so put a _ here
}
