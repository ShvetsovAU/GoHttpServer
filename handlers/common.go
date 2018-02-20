package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
	"time"
	"strconv"
	"encoding/json"
)

//имя параметра в контексте, содержащее id пользователя
const CtxUserId = "userId"

// Достать параметр из URL
func getParameter(r *http.Request, name string) string {
	params := context.Get(r, "params").(httprouter.Params)
	return params.ByName(name)
}

func getTimeParameter(r *http.Request, name string) time.Time {
	val, err := time.Parse(time.RFC3339, getParameter(r, name))
	if err != nil {
		panic(err)
	}
	return val
}

func getIntParameter(r *http.Request, name string, defaultValue int) int {
	val, err := strconv.Atoi(getParameter(r, name))
	if err != nil {
		val = defaultValue
	}
	return val
}

// Вернуть JSON-ответ клиенту
func writeResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(body)
}

// Вернуть JSON-ответ клиенту: Ok - если ошибок нет, иначе текст ошибки
func writeState(w http.ResponseWriter, err error) {
	if err == nil {
		writeResponse(w, "Ok")
	} else {
		panic(err)
	}
}

func writeBool(w http.ResponseWriter, b bool, err error) {
	if err == nil {
		writeResponse(w, b)
	} else {
		panic(err)
	}
}

func getUserId(r *http.Request) string {
	//м.б. стоит использовать "context.GetOk(r, constants.CtxUserId)" и обрабатывать ситуацию, когда параметр не найден?
	//return context.Get(r, constants.CtxUserId).(string)
	return context.Get(r, CtxUserId).(string)
}