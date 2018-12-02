package main

import (
	"html/template"
	"log"
	"net/http"
)

func displayJobStatus(responseWriter http.ResponseWriter, request *http.Request) {
	// fmt.Fprint(responseWriter, request.URL.Path)
	var jobStatusTemplate = template.New("job status")
	jobStatusTemplate, err := jobStatusTemplate.Parse(displayJobStatusTemplate)

	if err != nil {
		log.Println(err)
	}

	err = jobStatusTemplate.Execute(responseWriter, map[string]interface{}{
		"entries": cronObject.Entries(),
	})

	if err != nil {
		log.Println(err)
	}
}
