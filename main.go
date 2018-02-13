//package main
//
//import (
//	"fmt" // пакет для форматированного ввода вывода
//	"net/http" // пакет для поддержки HTTP протокола
//	//"strings" // пакет для работы с  UTF-8 строками
//	"log" // пакет для логирования
//	"strings"
//	//"github.com/shvetsovau/GoHttpServer/config"
//	"./config"
//	"io/ioutil"
//	"crypto/x509"
//	"crypto/tls"
//)
//
//func main() {
//
//	log.Println("Server starting...")
//
//	http.HandleFunc("/", HomeRouterHandler) // установим роутер
//
//	//ListenAndServeTLS Для HTPPS соеднинения
//	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
//	err := http.ListenAndServe(config.AppArgs.GetPort(), nil) // задаем слушать порт
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//
//	// Endoff
//	log.Println("Server is shut down")
//}
//

package main

import (
	"log" // пакет для логирования
	"github.com/shvetsovau/GoHttpServer/httpscerts"
	"github.com/shvetsovau/GoHttpServer/config"
	"net/http" // пакет для поддержки HTTP протокола
	"fmt" // пакет для форматированного ввода вывода
	"strings"
	"github.com/shvetsovau/GoHttpServer/router"
)

func RedirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Перенаправляем входящий HTTP запрос. Учтите,
	// что "127.0.0.1:8081" работает только для вашей локальной машине
	http.Redirect(w, r, config.AppArgs.GetFullHost()+r.RequestURI,
		http.StatusMovedPermanently)
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,
	fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Maksim!") // отправляем данные на клиентскую сторону
}

func AdminRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Хрен тебе!") // отправляем данные на клиентскую сторону
}

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Server starting...")

	// Проверяем, доступен ли cert файл.
	err := httpscerts.Check("cert.pem", "key.pem")

	// Если он недоступен, то генерируем новый.
	if err != nil {
		if config.CanCreateCertificatAndKey != false {
			err = httpscerts.Generate("cert.pem", "key.pem", config.AppArgs.GetFullHost())
			if err != nil {
				log.Fatal("Ошибка: Сервер не может сгенерировать https сертификат или приватный ключ.")
			}
		} else {
			log.Fatal("Ошибка: https сертификат или приватный ключ не найден.")
		}
	}

	// установим роутер
	//http.HandleFunc("/", HomeRouterHandler)
	//http.HandleFunc("/admin", AdminRouterHandler)
	// задаем слушать порт
	//err = http.ListenAndServeTLS(config.AppArgs.GetPort(), "cert.pem", "key.pem", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServeTLS: ", err)
	//}

	// Запуск HTTPS сервера (если нужно будет перенаправление с HTTP на HTTPS, нужно будет запускать HTTPS в отдельной go-рутине)
	http.ListenAndServeTLS(config.AppArgs.GetFullHost(), "cert.pem", "key.pem", router.GetRoutes())

	//// Запуск HTTP сервера и редирект всех входящих запросов на HTTPS
	//http.ListenAndServe(":8080", http.HandlerFunc(redirectToHttps))

	log.Println("Server is switch off")

	//cert, err := ioutil.ReadFile("server.crt")
	//if err != nil {
	//	log.Fatalf("Couldn't load file", err)
	//}
	//certPool := x509.NewCertPool()
	//certPool.AppendCertsFromPEM(cert)
	//
	//conf := &tls.Config{
	//	RootCAs: certPool,
	//}
	//
	//conn, err := tls.Dial("tcp", "localhost:443", conf)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer conn.Close()
	//
	//n, err := conn.Write([]byte("hello\n"))
	//if err != nil {
	//	log.Println(n, err)
	//	return
	//}
	//
	//buf := make([]byte, 100)
	//n, err = conn.Read(buf)
	//if err != nil {
	//	log.Println(n, err)
	//	return
	//}
	//
	//println(string(buf[:n]))
}