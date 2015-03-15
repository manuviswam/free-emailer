package email

import (
	"net/smtp"
)

type MailEngine struct {
	UserName string
	Password string
	EmailServer string
	Port string
}

func NewEngine(userName string,password string, server string, port string) *MailEngine {
	engine := new(MailEngine)
	engine.UserName = userName
	engine.Password = password
	engine.EmailServer = server
	engine.Port = port
	return engine
}

func (e MailEngine) Send(To string, Subject string, Content string) error {
	auth := smtp.PlainAuth("", e.UserName, e.Password, e.EmailServer)
	server := e.EmailServer + ":" + e.Port
	return smtp.SendMail(server, auth,e.UserName, []string{To}, getBytes(To, Subject, Content))
}

func getBytes(to string, sub string, cont string) []byte {
	return []byte("To:" + to + "\nSubject : " + sub + "\n" + cont)
}
