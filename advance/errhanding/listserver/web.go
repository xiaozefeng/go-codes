package main

import (
	"github.com/xiaozefeng/go-codes/advance/errhanding/listserver/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", errWrapper(handler.HandleFileList))

	log.Println("server start port 8888.")
	http.ListenAndServe(":8888", nil)
}

type UserError interface {
	error
	Message() string
}


// error handler
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// error wrapper
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handing request: %s\n", err.Error())
			if userErr, ok := err.(UserError);ok{
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
