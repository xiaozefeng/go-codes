package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		writer.Write(contents)
	})

	log.Println("server start port 8888.")
	http.ListenAndServe(":8888", nil)
}
