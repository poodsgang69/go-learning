package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Http Handler methods here
type Hello struct {
	l *log.Logger
}

// constructor for the Hello Struct
// anything starting with Caps can be exported to other packages.
// anything with lowercase is private to the current package.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	const ERROR_MSG string = "Bad Request"

	// log.Println("Response Handler for the '/helloworld' endpoint ")
	h.l.Println("Response Handler for the '/helloworld' endpoint ")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Option 1 using the http.Error
		// http.Error(rw, ERROR_MSG, http.StatusBadRequest)
		// Option 2 using the rw (io.writer)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(ERROR_MSG))
		return
	}

	log.Printf("Data in the Request: %s", data)

	// fmt.Fprintf(rw, "Hello %s\n", data) //one way to use the rw (io.writer) to directly inject data into it.
	var response string = fmt.Sprintf("Hello %s", data)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(response))
}
