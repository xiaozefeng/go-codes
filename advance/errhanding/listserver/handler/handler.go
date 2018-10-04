package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	if strings.Index(path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}

	filePath := path[len(prefix):]
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(contents)
	return nil
}
