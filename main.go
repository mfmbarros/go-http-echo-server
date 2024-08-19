package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const httpPort = "8080"

func main() {
	fmt.Printf("Starting http echo server on :%s\n", getPort())

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":"+getPort(), nil)

}

func getPort() string {
	var port string
	var ok bool
	if port, ok = os.LookupEnv("HTTP_ECHO_SERVER_PORT"); !ok {
		port = httpPort
	}
	return port
}

func amIHealthy() bool {
	return true
}

func healthHandler(writer http.ResponseWriter, request *http.Request) {
	if ok := amIHealthy(); ok {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, "HTTP Echo Server is healthy")
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, "HTTP Echo Server is unhealthy")
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

	request.Write(writer)
}
