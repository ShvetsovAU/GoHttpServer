package utils

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

/*
	Если есть ошибка - паниковать. Кроме того, в лог будет записано место возникновения ошибки.
	К ошибке можно добавить метку (или несколько через запятую)
	Пример: вызов PanicIfErr(error["not found"], "UserBl") приведет к панике с ошибкой: "[UserBl]: not found"
*/
func PanicIfErr(err error, tags ...string) {
	if err == nil {
		return
	}

	err = CreateTaggedErr(err, tags...)

	// мб стоит это оформить как "if Debug...."
	_, file, line, _ := runtime.Caller(1)
	log.Printf("Debug info for %v. Details: [%v], line [%v]", err, file, line)

	panic(err)
}

/*
	Добавляет метку (или несколько через запятую) к тексту ошибки
	Пример: вместо ошибки "not found" можно получить более осмысленное "[device]: not found"
*/
func CreateTaggedErr(err interface{}, tags ...string) error {
	if err == nil {
		return nil
	}

	if len(tags) == 0 {
		return fmt.Errorf("%v", err)
	}

	return fmt.Errorf("[%v]: %v", strings.Join(tags, ","), err)
}

func ErrorText(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()

}