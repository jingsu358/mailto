package sendmail

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	//	"strings"

	"github.com/kylelemons/go-gypsy/yaml"
)

func Readyaml() []string {
	config, err := yaml.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	code := []string{}
	//get mail, password  and other info
	sendmail, _ := config.Get("mail")
	password, _ := config.Get("password")
	smtpserver, _ := config.Get("smtp")
	port, _ := config.Get("port")
	content_type, _ := config.Get("content_type")
	//	fmt.Printf("v2 type:%T\n", port)
	code = append(code, sendmail, password, smtpserver, port, content_type)
	//	return sendmail, password, smtpserver, port, content_type
	return code
}

func Sendmail(receive []string, m map[string]string) {
	code := Readyaml()
	//	mail, password, smtpserver, port, content_type := Readyaml()
	mail, password, smtpserver, port, content_type := code[0], code[1], code[2], code[3], code[4]
	//	mail = mail[1 : len(mail)-1]
	auth := smtp.PlainAuth("", mail, password, smtpserver)
	to := receive
	name := m
	//	fmt.Println(m[])
	//	nickname := "nil"
	user := mail
	subject := "Happy birthday to you!"
	// get mail body content
	filebody, bodyerr := ioutil.ReadFile("body.txt")
	if bodyerr != nil {
		fmt.Printf("Open file body.txt Error: %s\n", bodyerr)
		panic(bodyerr)
	}
	for _, line := range to {
		body := "Dear " + name[line] + ":\n" + string(filebody)
		//once send to one person
		msg := []byte("To:" + line + "\r\nFrom: " + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
		// receive args must be [] , string will error
		err := smtp.SendMail(smtpserver+":"+port, auth, user, []string{line}, msg)
		if err != nil {
			fmt.Printf("send mail error: %v", err)
		} else {
			fmt.Printf("send mail to %v seccess! \n", line)
		}

	}
	//	fmt.Println(body)
	//	msg := []byte("To:" + strings.Join(to, ",") + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	//	msg := []byte("To:" + strings.Join(to, ",") + "\r\nFrom: " + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	//	//	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	//	err := smtp.SendMail(smtpserver+":"+port, auth, user, to, msg)
	//	if err != nil {
	//		fmt.Printf("send mail error: %v", err)
	//	} else {
	//		fmt.Printf("send mail to %v seccess! \n", to)
	//	}
}
