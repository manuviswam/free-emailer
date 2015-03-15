package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/",Index)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func Index(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "I am alive!!!")
}
