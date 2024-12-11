//go:build file1
// +build file1

package main

import (
	"context"
	"log"
	"module/new/directory/API/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const ERROR_MSG string = "Bad Request"

// Implementation of the below defined goroutine calling a func
// func startServer() {
// 	var customServeErr error = customServer.ListenAndServe()
// 	if customServeErr != nil {
// 		l.Fatal("Server Error: ", customServeErr)
// 	}
// }

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
	helloHandler := handlers.NewHello(l)
	helloDefaultHandler := handlers.NewHelloDefault(l)
	productHandler := handlers.NewProduct(l)

	overallServeMux := http.NewServeMux()
	overallServeMux.Handle("/", helloDefaultHandler)
	overallServeMux.Handle("/helloworld", helloHandler)
	overallServeMux.Handle("/products", productHandler)

	/*
		Creating a Server to more flexibily configure server properties like ReadTime WriteTime etc.
	*/
	var customServer *http.Server = &http.Server{
		Addr:         ":9090",
		Handler:      overallServeMux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	/*
		this will block the below code.
		hence we define a GoRoutine so that it funcs in a nonblocking way for the shutdown code below.

		Also there are 2 ways to define a goRoutine Function. this is an inline go func
	*/
	go func() {
		l.Println("Starting Server on port: 9090")
		var customServeErr error = customServer.ListenAndServe()
		if customServeErr != nil {
			l.Fatal("Server Error: ", customServeErr)
		}
	}()

	/*
		This is a go routine calling a function
	*/
	// go startServer()

	/*
		creating a channel for os.Signal type.
		channels allow for data transfer and communication between threads
	*/
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	//getting the value sent by the signalChannel to the signalConsumer variable
	signalConsumer := <-signalChannel
	l.Println("Recieved terminate signal, Gracefully Exiting", signalConsumer)

	/*
		Handling Graceful exits: using the shutdown parameter, we can handle server shutdowns without modifying or disturbing ongoing requests.
		We need to define a context Deadline specifying under which context the server shutdown needs to be initiated.

		Contexts are used for graceful exits and also passing params such as api keys and cookies, etc.
		Search for examples to understand the use of contexts, go routines and defer.

		context.WithDeadline() is used to tell the context to exit after certain time. For this we need to know the exact time (like a cron -> needs to run at midnight or at this time).
		context.Background() is the default root context.
		context.WithTimeout is used when we define custom relative timeouts.
	*/
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	customServer.Shutdown(timeoutContext)

	// http.ListenAndServe(":9090", http.DefaultServeMux)
	// http.ListenAndServe(":9090", overallServeMux)

}

/*
	Separate function outside func
*/
// func defaultPathHandler(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("Response Handler for the default url / ")
// }
