package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Experience    experience    `json:"experience"`
	Salary        int           `json:"salary"`        // 40000
	Qualification qualification `json:"qualification"` //bachelor in CSIT
	Location      string        `json:"location"`      //  Pokhara
	JobTitle      string        `json:"title"`         //  Accountant
	JobCategory   string        `json:"category"`      // IT and Telecommunications
	JobType       string        `json:"type"`          // full time , part time
	JobLevel      string        `json:"level"`         // senior, junior, mid
}

type experience struct {
	Value int    `json:"value"`
	Type  string `json:"type`
}
type qualification struct {
	Value string `json:"value"` // bachelor ,master
	Type  string `json:"type`   //bsc csit,  Computer engineering
}

var searchRequest Request

// Handles the post Request to naviai-service.
func RequestHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		log.Println("new POST request arrived", req)
		err := json.NewDecoder(req.Body).Decode(&searchRequest)
		if err != nil {
			log.Println("error in decoding request ", err)
			exp := experience{
				Value: 2,
				Type:  "years",
			}
			qual := qualification{
				Type:  "CSIT",
				Value: "bachelor",
			}
			searchRequest.Experience = exp
			searchRequest.JobCategory = "Finance"
			searchRequest.JobType = "full time "
			searchRequest.JobLevel = "mid"
			searchRequest.JobTitle = "Accountant"
			searchRequest.Location = "any"
			searchRequest.Salary = 20000
			searchRequest.Qualification = qual

			processSearch()
		} else {
			// normal processing
			log.Println("normal processing", searchRequest)
			processSearch()
		}
	}
}
