package handler

import (
	"fmt"
	"log"

	"github.com/nabinkh/go-backend/fetch"
	"github.com/nabinkh/go-backend/parser"
)

type SiteInfo struct {
	SearchURL            string
	SearchByCategory     string
	SearchByJobTitle     string
	SearchByOrganization string
}

var sites []string

func init() {
	sites = []string{
		"merojob",
	}
}
func processSearch() {
	log.Println("processing the info for all the available sites ..... ")
	for _, site := range sites {
		info := getURL(site, searchRequest.JobTitle)
		if searchRequest.JobTitle != "any" {
			fetch.MakeGet(info.SearchURL)
		} else if searchRequest.JobCategory != "any" {
			pageSource := fetch.MakeGet(info.SearchByCategory)
			parser.ParseCategoryPage(pageSource, site)

		}
	}
}

func getURL(site string, searchString string) SiteInfo {
	var URL SiteInfo
	switch site {
	case "merojob":
		URL.SearchURL = fmt.Sprintf("https://merojob.com/?q=%s", searchString)
		URL.SearchByCategory = fmt.Sprintf("https://merojob.com/category/")
		URL.SearchByJobTitle = fmt.Sprintf("https://merojob.com/designation")
		URL.SearchByOrganization = fmt.Sprintf("https://merojob.com/company")
	}
	return URL
}
