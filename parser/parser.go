package parser

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Request struct {
	Experience    string `json:"experience"`
	Salary        string `json:"salary"`        // 40000
	Qualification string `json:"qualification"` //bachelor in CSIT
	Location      string `json:"location"`      //  Pokhara
	JobTitle      string `json:"title"`         //  Accountant
	JobCategory   string `json:"category"`      // IT and Telecommunications
	JobType       string `json:"type"`          // full time , part time
	JobLevel      string `json:"level"`         // senior, junior, mid
}

var retrivedData Request

//ParseCategoryPage ...
func ParseCategoryPage(pageSource string, site string) []string {
	var arr []string
	return arr
}

// ParseSearchPage ...
func ParseSearchPage(pageSource string, site string) []string {
	var resultArray []string
	p := strings.NewReader(pageSource)
	doc, _ := goquery.NewDocumentFromReader(p)
	doc.Find(".card-block").Each(func(i int, s *goquery.Selection) {
		resultUrl, found := s.Find(".col-lg-9").Find("h1").Find("a").Attr("href")
		if found {
			resultArray = append(resultArray, fmt.Sprintf("http://%s.com%s", site, resultUrl))
		}
	})
	return resultArray
}

func ParseVacancyPage(pgsrc string, site string) {
	p := strings.NewReader(pgsrc)
	doc, _ := goquery.NewDocumentFromReader(p)
	doc.Find(".card-group").Each(func(i int, s *goquery.Selection) {
		_, exists := s.Find(".card-header").Find("h3").Attr("class")
		if exists {
			header := s.Find(".card-header").Find("h3")
			if strings.Contains(header.Text(), "Specification") {
				// parse the inside of the specification for particular website
				s.Find(".card-block").Find("table").Find("tr").Each(func(i int, s *goquery.Selection) {
					s.Children().Each(func(j int, t *goquery.Selection) {
						// fmt.Println(j, t.Text())
						if j == 0 {
							saveData(t)
						}
					})
				})
			} else if strings.Contains(header.Text(), "Information") {
				// parse the inside of information for a particular website
				s.Find(".card-block").Find("table").Find("tr").Each(func(i int, s *goquery.Selection) {
					s.Children().Each(func(j int, t *goquery.Selection) {
						// fmt.Println(j, t.Text())
						if j == 0 {
							saveData(t)
						}
					})
				})
			}
		}
	})
	fmt.Println("the data retrived from website is", retrivedData)
}
func saveData(t *goquery.Selection) {
	nextVal := strings.Trim(t.Next().Text(), " ")
	// save data to request object
	if strings.Contains(strings.ToLower(t.Text()), "category") {
		retrivedData.JobCategory = nextVal
	} else if strings.Contains(strings.ToLower(t.Text()), "level") {
		retrivedData.JobLevel = nextVal
	} else if strings.Contains(strings.ToLower(t.Text()), "location") {
		retrivedData.Location = nextVal
	} else if strings.Contains(strings.ToLower(t.Text()), "salary") {
		retrivedData.Salary = nextVal
	} else if strings.Contains(strings.ToLower(t.Text()), "education") {
		retrivedData.Qualification = nextVal
	} else if strings.Contains(strings.ToLower(t.Text()), "experience") {
		retrivedData.Experience = nextVal
	}
}
