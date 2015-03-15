package main

import (
	"net/http"
	"log"
	"config"
	"email"
	"fmt"
)

var sender *email.MailEngine

func main(){
	cfg,_ := config.LoadConfig("config.gcfg")

	fileServer := http.FileServer(http.Dir(cfg.Http.StaticDirPath))
	mailServer := cfg.MailServer
	sender = email.NewEngine(mailServer.UserName, mailServer.Password, mailServer.ServerUrl, mailServer.MailPort)

	http.Handle("/",fileServer)
	http.HandleFunc("/mail", sendMail)
	log.Fatal(http.ListenAndServe(cfg.Http.Port,nil))
}

func sendMail(writer http.ResponseWriter, request *http.Request) {
	err := sender.Send(request.PostFormValue("To"),request.PostFormValue("Subject"),request.PostFormValue("Content"))
	fmt.Fprintf(writer, "error:%s\nTo:%s\nSub:%s",err,request.PostFormValue("To"),request.PostFormValue("Subject"))
}
