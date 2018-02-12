package config

import (
	"fmt"
	"log"
	"github.com/jessevdk/go-flags"
	"os"
)

const (
	portMin          = 8000
	portMax          = 8099
	portDefault      = 8083
	hostDefault      = "localhost"
	errUserParse     = "[EXIT] User/Password not specified (-u user:password)"
	errPortIncorrect = "Error: Port should be in a range [%v:%v]"
	msgPortSetDefault = "Set default port value: %v"
)

//Может ли сервер самостоятельно генерировать сертифкат и приватный ключ
var CanCreateCertificatAndKey bool

var AppArgs *Args

type Args struct {
	Port       uint   `short:"p" long:"port" description:"Server port" default:"8083"`
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
	AppArgs.Info()
}

func (this *Args) init() {
	_, err := flags.ParseArgs(AppArgs, os.Args)

	if err != nil {
		log.Fatal(err)
	}

	if this.Port < portMin || this.Port > portMax {
		//log.Fatal(fmt.Sprintf(errPortIncorrect, portMin, portMax))
		log.Println(fmt.Sprintf(errPortIncorrect, portMin, portMax))
		log.Println(fmt.Sprintf(msgPortSetDefault, portDefault))
		//SetDefaultPort()
		SetDefaultCertificateValue()
	}

	//if up := strings.Split(this.User, ":"); len(up) == 2 {
	//	this.User = up[0]
	//	this.Password = up[1]
	//} else {
	//	log.Fatal(errUserParse)
	//}
}

func (this *Args) Info() {
	log.Printf("Command line arguments port:%v host:%v db:%v u:%v/%v qhost:%v qmail:%v jobs:%v\n",
		this.Port, this.Host, this.DB, this.User, this.Password, this.AmqpHost, this.MailsQueue, this.JobsCfg)
}

//Получить порт из настроек
func (this *Args) GetPort() string {

	//Если строка пустая, устанавливаем значение по умолчанию
	if this.Port == 0 {
		SetDefaultPort()
	}

	return fmt.Sprintf(":%v", this.Port)
}

//Получить хост из настроек
func (this *Args) GetHost() string {

	//Если строка пустая, устанавливаем значение по умолчанию
	if len(this.Host) == 0 {
		SetDefaultHost()
	}

	return fmt.Sprintf("%v", this.Host)
}

//Получить полный пусть к хосту (хост + порт)
func (this *Args) GetFullHost() string {
	return fmt.Sprintf("%v%v", this.GetHost(), this.GetPort())
}

//Установить порт, используемый по умолчанию
func SetDefaultPort() {
	AppArgs.Port = portDefault
}

//Установить хост, используемый по умолчанию
func SetDefaultHost() {
	AppArgs.Host = hostDefault
}

//Установить значение, разрешающее серверу самостоятельно генерировать сертификат и ключ (пока для написания сервера,
// а потом видно будет)
func SetDefaultCertificateValue() {
	CanCreateCertificatAndKey = true
}