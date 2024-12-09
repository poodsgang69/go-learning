package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloDefault struct {
	l *log.Logger
}

func NewHelloDefault(l *log.Logger) *HelloDefault {
	return &HelloDefault{l}
}

func (h *HelloDefault) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	const ERROR_MSG string = "Bad Request"

	h.l.Println("Default Handler for URL /")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, ERROR_MSG, http.StatusBadRequest)
		return
	}

	h.l.Printf("Data in the Request is: %s\n", data)

	var response string = fmt.Sprintf("\n(/) Data: %s", data)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(response))
}
