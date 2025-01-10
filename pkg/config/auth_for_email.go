package config

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

func (m *AppConfig) Create_Email_auth(id string, server string, port int64, account string, pwd string, check bool) (re_err error) {
	defer func() { //this function recovers from panic
		r := recover()
		if r != nil {
			fmt.Println(r)
			re_err = fmt.Errorf("%v", r)
		}
	}()

	if pwd == "" {
		fmt.Println("Please provide the pasword for:", account, "on server", server)
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		pwd = line
	}
	auth := smtp.PlainAuth(id, account, pwd, server)

	var data Email_sending
	data.Auth = auth
	data.Addr = server + ":" + strconv.FormatInt(port, 10)
	data.Sender = account
	m.Email_auth[id] = data

	if check {
		//we try to send a testing email, in order to make sure everything works
		//TODO implement
	}

	return nil
}
