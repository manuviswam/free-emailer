package main

import (
	"net/http"
	"log"
	"config"
)

func main(){
	fileServer := http.FileServer(http.Dir("static_files"))
	http.Handle("/",fileServer)
	cfg,_ := config.LoadConfig("config.gcfg")
	log.Fatal(http.ListenAndServe(cfg.Http.Port,nil))
}
