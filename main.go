package main

import (
	"fmt" // пакет для форматированного ввода вывода
	"net/http" // пакет для поддержки HTTP протокола
	//"strings" // пакет для работы с  UTF-8 строками
	"log" // пакет для логирования
	"strings"
	//"config"
)

func main() {

	log.Println("Server starting...")

	http.HandleFunc("/", HomeRouterHandler) // установим роутер

	//ListenAndServeTLS Для HTPPS соеднинения
	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
	err := http.ListenAndServe(config.AppArgs.GetPort(), nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// Endoff
	log.Println("Server is shut down")
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