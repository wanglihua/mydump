package main

import (
	"fmt"
	"net/http"
)

func displayJobStatus(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, request.URL.Path)
}
