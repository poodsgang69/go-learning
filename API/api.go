package main

import (
	"log"
	"net/http"
)

func main2() {

	/*
		ANONYMOUS FUNCTION
	*/
	defaultPathHandler := func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Response Handler for the default url / ")
	}

	http.HandleFunc("/", defaultPathHandler)
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Response Handler url /")
	// })

	// http.ListenAndServe(":9090", http.DefaultServeMux)
	http.ListenAndServe(":9090", nil)

}

/*
	Separate function outside func
*/
// func defaultPathHandler(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("Response Handler for the default url / ")
// }
