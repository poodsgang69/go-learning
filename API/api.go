//go:build file1
// +build file1

package main

import (
	"log"
	"module/new/directory/API/handlers"
	"net/http"
	"os"
)

const ERROR_MSG string = "Bad Request"

func main() {

	// /*
	// 	ANONYMOUS FUNCTION
	// */
	// defaultPathHandler := func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Response Handler for the default url / ")
	// }

	// helloWorldPathHandler := func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Response Handler for the '/helloworld' endpoint ")
	// 	data, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		// Option 1 using the http.Error
	// 		// http.Error(rw, ERROR_MSG, http.StatusBadRequest)
	// 		// Option 2 using the rw (io.writer)
	// 		rw.WriteHeader(http.StatusBadRequest)
	// 		rw.Write([]byte(ERROR_MSG))
	// 		return
	// 	}

	// 	log.Printf("Data in the Request: %s", data)

	// 	// fmt.Fprintf(rw, "Hello %s\n", data) //one way to use the rw (io.writer) to directly inject data into it.
	// 	var response string = fmt.Sprintf("Hello %s", data)
	// 	rw.WriteHeader(http.StatusOK)
	// 	rw.Write([]byte(response))
	// }

	// http.HandleFunc("/", defaultPathHandler)
	// // http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// // 	log.Println("Response Handler url /")
	// // })

	// http.HandleFunc("/helloworld", helloWorldPathHandler)

	/*
		We are registering a custom serveHTTP method, which overrides the defaultServeMux
	*/
	l := log.New(os.Stdout, "API-GO Logger: ", log.LstdFlags)
	hh := handlers.NewHello(l)

	dsm := http.NewServeMux()
	dsm.Handle("/helloworld", hh)

	// http.ListenAndServe(":9090", http.DefaultServeMux)
	http.ListenAndServe(":9090", dsm)

}

/*
	Separate function outside func
*/
// func defaultPathHandler(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("Response Handler for the default url / ")
// }
