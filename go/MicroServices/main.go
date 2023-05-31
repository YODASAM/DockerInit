package main

import (
	"log"
	"net/http"
)

func main() { // main entry point for Go Code

	// Create Webserver : ? How - We need a HTTP Package

	// HTTP Package has a HUGEEEE AMOUNT OF  capability for building Web Services

	//1. Basice step of creating is Http ListenandServe - takes an address and terh handler
	// First a bind address
	// I want to bind every Ip address on my machine to port 9090,
	// for specific ipaddress I could give 128.1.1..15: 9090 etc
	//http.ListenAndServe(":9090",nil)

	// STEP 2 : we create a http handler to handle requests:
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World!! Sam")
	})

	// The above handler creates an interface for the function
	// and it adds it to the default Serve MUX

	// Handlefunc is a convenient method  in Go HTTP package which
	//  It registers a function to a path on Default Serve  MUX
	// DEfault MUX is a Http Handler
	// And everything related to the server in go is an HTTP Handler

	// Lets add another handler - goodbye
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Tata World!! Sam")
		// when the curl path matches godobye - this Tata function comment is called
		// anything other than goodbye - default handler is called - Default serve MUX
	})

	// STEP 2 was just Loggign for different paths
	// Now STEP 3: we will see how to read and write :
	// we could do that by using a ResponseWriter and a HTTP Request

	// To construct a HTTP response -
	// methods in reponseWriter:
	// Writing heaedrs, response bodies, status codes etc
	// concrete input is injected by the go server durign execution

	// in STEP 3 - first we want to read from teh body :
	https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_
	27:15

	// STEP 1: This line create the server
	http.ListenAndServe(":9090", nil)
	// This convenience method constructs an HTTP Server and registers a
	// default handler to it
	//  if we do not specify any handler e.g. nil , it takes the efault serve MUX

	// sERVE mux IS ACTUALLY A REQUEST MultiPlexer
	// HAndle with register a Handler for th e path
	// So Handle will register an interface for a given path
	// Handler in go is just an interface

}
