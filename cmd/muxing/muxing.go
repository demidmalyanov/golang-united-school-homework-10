package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func paramHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["param"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func badReqHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func bodyDataHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "I got message:\n%s", msg)
}

func headersDataHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		log.Println(err)
	}

	sum := strconv.Itoa(a + b)
	w.Header().Set("a+b", sum)
}

func defaultHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", paramHandler)
	router.HandleFunc("/bad", badReqHandler)
	router.HandleFunc("/data", bodyDataHandler).Methods("POST")
	router.HandleFunc("/headers", headersDataHandler).Methods("POST")
	router.HandleFunc("/",defaultHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
