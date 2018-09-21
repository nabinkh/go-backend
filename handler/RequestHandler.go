package handler

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/nabinkh/go-backend/parser"
)

var searchRequest Request

// Handles the post Request to naviai-service.
func RequestHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// log.Println("new POST request arrived", req)
		err := json.NewDecoder(req.Body).Decode(&searchRequest)
		if err != nil {
			log.Println("error in decoding request ", err)
			searchRequest.Experience = "2 years"
			searchRequest.JobCategory = "Finance"
			searchRequest.JobType = "full time "
			searchRequest.JobLevel = "mid"
			searchRequest.JobTitle = "Accountant"
			searchRequest.Location = "any"
			searchRequest.Salary = "20000"
			searchRequest.Qualification = "bachelor"

			processSearch()
		} else {
			// normal processing
			log.Println("normal processing", searchRequest)
			processSearch()
		}
	}
}
