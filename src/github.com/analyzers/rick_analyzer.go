package analyzers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type RickAnalyzer struct {
	URL string
}

func (analyzer RickAnalyzer) makeRequest(url string) map[string]interface{} {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(responseData)
	var decodedResponse interface{}
	json.Unmarshal(bytes, &decodedResponse)
	return decodedResponse.(map[string]interface{})
}

func (analyzer RickAnalyzer) getInformation(url string) (map[string]interface{}, []string) {
	response := analyzer.makeRequest(url)
	info := response["info"].(map[string]interface{})
	results := response["results"].([]interface{})
	var names []string
	for i := range results {
		result := results[i].(map[string]interface{})
		names = append(names, strings.ToLower(result["name"].(string)))
	}
	return info, names
}

func (analyzer RickAnalyzer) getNames(url string, names chan string) {
	response := analyzer.makeRequest(url)
	results := response["results"].([]interface{})

	for i := range results {
		result := results[i].(map[string]interface{})
		names <- strings.ToLower(result["name"].(string))
	}
}

func (analyzer RickAnalyzer) GetAllNames() []string {
	urlBase := analyzer.URL
	firstPageURL := urlBase + "?page=1"
	infoLocations, locations := analyzer.getInformation(firstPageURL)
	numPages := int(infoLocations["pages"].(float64))
	totalLocations := int(infoLocations["count"].(float64)) - len(locations)
	names := make(chan string)

	for i := 2; i <= numPages; i++ {
		url := urlBase + "?page=" + strconv.Itoa(i)
		go analyzer.getNames(url, names)
	}

	for i := 1; i <= totalLocations; i++ {
		locations = append(locations, <-names)
	}

	return locations
}

func (analyzer RickAnalyzer) CountLetters(names []string, letter string) int {
	num := 0
	for _, name := range names {
		num = num + strings.Count(name, letter)
	}
	return num
}
