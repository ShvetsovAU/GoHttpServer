package handlers

import (
	"net/http"
	"fmt"
	"strings"
)

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
	fmt.Fprintf(w, "Hello User!") // отправляем данные на клиентскую сторону
}

func AdminRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Хрен тебе!") // отправляем данные на клиентскую сторону
}