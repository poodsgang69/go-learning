//go:build file1
// +build file1

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	/*
		ANONYMOUS FUNCTION
	*/
	defaultPathHandler := func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Response Handler for the default url / ")
	}

	helloWorldPathHandler := func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Response Handler for the '/helloworld' endpoint ")
		data, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data in the Request: %s", data)

		// fmt.Fprintf(rw, "Hello %s\n", data) //one way to use the rw (io.writer) to directly inject data into it.
		var response string = fmt.Sprintf("Hello %s", data)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(response))
	}

	http.HandleFunc("/", defaultPathHandler)
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Response Handler url /")
	// })

	http.HandleFunc("/helloworld", helloWorldPathHandler)

	// http.ListenAndServe(":9090", http.DefaultServeMux)
	http.ListenAndServe(":9090", nil)

}

/*
	Separate function outside func
*/
// func defaultPathHandler(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("Response Handler for the default url / ")
// }
