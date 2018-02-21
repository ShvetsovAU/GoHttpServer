package main

import (
	"log" // пакет для логирования
	"github.com/shvetsovau/GoHttpServer/httpscerts"
	"github.com/shvetsovau/GoHttpServer/config"
	"net/http" // пакет для поддержки HTTP протокола
	//"fmt" // пакет для форматированного ввода вывода
	"github.com/shvetsovau/GoHttpServer/router"
)

func RedirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Перенаправляем входящий HTTP запрос. Учтите,
	// что "127.0.0.1:8081" работает только для вашей локальной машине
	http.Redirect(w, r, config.AppArgs.GetFullHost()+r.RequestURI,
		http.StatusMovedPermanently)
}

func main() {
	log.SetFlags(log.Lshortfile)

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

	log.Println("Server starting...")

	// Запуск HTTPS сервера (если нужно будет перенаправление с HTTP на HTTPS, нужно будет запускать HTTPS в отдельной go-рутине)
	http.ListenAndServeTLS(config.AppArgs.GetFullHost(), "cert.pem", "key.pem", router.GetRoutes())

	//// Запуск HTTP сервера и редирект всех входящих запросов на HTTPS
	//http.ListenAndServe(config.AppArgs.GetFullHost(), router.GetRoutes())

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