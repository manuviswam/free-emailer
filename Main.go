package main

import (
	"fmt"
	"net/http"
	"log"
	"config"
)

func main(){
	http.HandleFunc("/",Index)
	cfg,_ := config.LoadConfig("config.gcfg")
	log.Fatal(http.ListenAndServe(cfg.Http.Port,nil))
}

func Index(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "I am alive!!!")
}
