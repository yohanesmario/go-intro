package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/html")

	io.WriteString(res, "<div style='font-size:200px'>TEST</div>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	fmt.Println("listening on localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}
