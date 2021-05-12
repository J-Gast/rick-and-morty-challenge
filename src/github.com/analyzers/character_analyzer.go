package analyzers

import (
	"rick-and-morty-challenge/src/github.com/models"
	"strconv"
)

type CharactersAnalyzer struct {
	analyzer      RickAnalyzer
	url           string
	CharactersMap map[string]string
}

func (character *CharactersAnalyzer) Init() {
	character.url = "https://rickandmortyapi.com/api/character"
	character.analyzer = RickAnalyzer{URL: character.url}
}

func (character CharactersAnalyzer) GetOriginNames(url string, page string, channel chan models.Character) {
	response := character.analyzer.MakeRequest(url + "?page=" + page)
	results := response["results"].([]interface{})
	for i := range results {
		characterModel := models.Character{}
		result := results[i].(map[string]interface{})
		id := strconv.Itoa(int(result["id"].(float64)))
		characterModel.URL = url + "/" + id
		characterModel.Origin = result["origin"].(map[string]interface{})["name"].(string)
		channel <- characterModel
	}
}

func (character *CharactersAnalyzer) GetAll() {
	baseURL := "https://rickandmortyapi.com/api/character"
	character.CharactersMap = make(map[string]string)
	info, _ := character.analyzer.GetInformation(baseURL)
	numPages := int(info["pages"].(float64))
	totalCharacters := int(info["count"].(float64))

	channel := make(chan models.Character)

	for i := 1; i <= numPages; i++ {
		page := strconv.Itoa(i)
		go character.GetOriginNames(baseURL, page, channel)
	}

	for i := 1; i < totalCharacters; i++ {
		characterModel := <-channel
		character.CharactersMap[characterModel.URL] = characterModel.Origin
	}
}

func (character CharactersAnalyzer) GetAllNames(channel chan string) {
	names := character.analyzer.GetAllNames()
	for i := range names {
		channel <- names[i]
	}
	close(channel)
}

func (character CharactersAnalyzer) CountLetters(names []string, channel chan int) {
	count := character.analyzer.CountLetters(names, "c")
	channel <- count
	close(channel)
}
