package fetch

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var httpClient http.Client

func init() {
	httpClient = http.Client{Timeout: time.Second * 5}
}

func MakeGet(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error creating request ", err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// log.Println("\n response from site is \n\n", string(body))
	return string(body)
}
