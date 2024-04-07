package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	fmt.Println("Server is listening port 3002...")
	http.ListenAndServe(":3002", nil)
}
