package handlers

import (
	"log"
	"module/new/directory/API/commons"
)

// type Hello struct {
// 	l *log.Logger
// }

var HelloObj commons.Hello = &commons.Hello

func NewHello(l *log.Logger) *HelloObj {
	return &HelloObj{l}
}

func serveHTTP() {

}
