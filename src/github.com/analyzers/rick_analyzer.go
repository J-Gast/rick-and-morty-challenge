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

func (analyzer RickAnalyzer) MakeRequest(url string) map[string]interface{} {
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

func (analyzer RickAnalyzer) Analyze() []int {
	listOfAnalyzers := make([]Analyzer, 3)
	listOfAnalyzers[0] = &LocationsAnalyzer{}
	listOfAnalyzers[1] = &CharactersAnalyzer{}
	listOfAnalyzers[2] = &EpisodeAnalyzer{}

	var listOfChannels [3]chan string
	for i := range listOfChannels {
		listOfChannels[i] = make(chan string)
	}

	for i := 0; i < 3; i++ {
		listOfAnalyzers[i].Init()
		go listOfAnalyzers[i].GetAllNames(listOfChannels[i])
	}

	var locationsNames []string
	var charactersNames []string
	var episodeNames []string

	for name := range listOfChannels[0] {
		locationsNames = append(locationsNames, name)
	}
	for name := range listOfChannels[1] {
		charactersNames = append(charactersNames, name)
	}
	for name := range listOfChannels[2] {
		episodeNames = append(episodeNames, name)
	}

	var listOfNames [3][]string
	listOfNames[0] = locationsNames
	listOfNames[1] = charactersNames
	listOfNames[2] = episodeNames

	var listOfCountChannels [3]chan int
	for i := range listOfChannels {
		listOfCountChannels[i] = make(chan int)
	}

	for i := 0; i < 3; i++ {
		go listOfAnalyzers[i].CountLetters(listOfNames[i], listOfCountChannels[i])
	}

	list := make([]int, 3)
	for i := range listOfChannels {
		list[i] = <-listOfCountChannels[i]
	}
	return list
}

func (analyzer RickAnalyzer) GetInformation(url string) (map[string]interface{}, []string) {
	response := analyzer.MakeRequest(url)
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
	response := analyzer.MakeRequest(url)
	results := response["results"].([]interface{})

	for i := range results {
		result := results[i].(map[string]interface{})
		names <- strings.ToLower(result["name"].(string))
	}
}

func (analyzer RickAnalyzer) GetAllNames() []string {
	urlBase := analyzer.URL
	firstPageURL := urlBase + "?page=1"
	infoLocations, locations := analyzer.GetInformation(firstPageURL)
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
