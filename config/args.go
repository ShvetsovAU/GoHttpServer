package config

import (
	"fmt"
	"os"
	"log"
	"github.com/jessevdk/go-flags"
)

const (
	portMin          = 8000
	portMax          = 8099
	errUserParse     = "[EXIT] User/Password not specified (-u user:password)"
	errPortIncorrect = "Error: Port should be in a range [%v:%v]"
)

var AppArgs *Args

type Args struct {
	Port       uint   `short:"p" long:"port" description:"Server port" default:"8082"`
	Host       string `short:"a" long:"address" description:"Hostname or address" default:"localhost"`
	DB         string `short:"d" long:"database" description:"Database name" default:"bills"`
	User       string `short:"u" long:"user" description:"user:password" default:"user:qwerty"`
	Password   string
	AmqpHost   string      `short:"q" long:"qhost" description:"RabbitMQ host" default:"amqp://"`
	MailsQueue string      `short:"m" long:"qmail" description:"RabbitMQ mails queue" default:"MailsQueue"`
	JobsCfg    string      `short:"j" long:"jobs" description:"Jobs config name" default:"jobs.cfg"`
	T          interface{} `short:"t"` // чтобы тесты не падали
}

func init() {
	AppArgs = new(Args)
	AppArgs.init()
	//AppArgs.Info()
}

func (this *Args) init() {
	_, err := flags.ParseArgs(AppArgs, os.Args)

	if err != nil {
		log.Fatal(err)
	}

	if this.Port < portMin || this.Port > portMax {
		log.Fatal(fmt.Sprintf(errPortIncorrect, portMin, portMax))
	}

	//if up := strings.Split(this.User, ":"); len(up) == 2 {
	//	this.User = up[0]
	//	this.Password = up[1]
	//} else {
	//	log.Fatal(errUserParse)
	//}
}

func (this *Args) GetPort() string {
	return fmt.Sprintf(":%v", this.Port)
}