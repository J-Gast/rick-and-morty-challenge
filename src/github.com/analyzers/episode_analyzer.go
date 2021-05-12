package analyzers

import (
	"rick-and-morty-challenge/src/github.com/models"
	"strconv"
)

type EpisodeAnalyzer struct {
	analyzer          RickAnalyzer
	characterAnalyzer CharactersAnalyzer
	url               string
}

func (episode *EpisodeAnalyzer) Init() {
	episode.url = "https://rickandmortyapi.com/api/episode"
	episode.analyzer = RickAnalyzer{URL: episode.url}
	episode.characterAnalyzer = CharactersAnalyzer{}
}

func (episode EpisodeAnalyzer) getOriginsByEachEpisode(url string, page string, channel chan models.Episode) {
	response := episode.analyzer.MakeRequest(url + "/" + page)

	charactersURLS := response["characters"].([]interface{})
	episodeID := int(response["id"].(float64))
	episodeName := response["episode"].(string)
	episodeModel := models.Episode{}
	episodeModel.Init()
	episodeModel.ID = episodeID
	episodeModel.EpisodeName = episodeName
	for i := range charactersURLS {
		URL := charactersURLS[i].(string)
		episodeModel.OriginLocations[episode.characterAnalyzer.CharactersMap[URL]] = true
	}
	channel <- episodeModel
}

func (episode EpisodeAnalyzer) GetOrigins() []models.Episode {
	episode.characterAnalyzer.GetAll()
	baseURL := "https://rickandmortyapi.com/api/episode"
	info, _ := episode.analyzer.GetInformation(baseURL)
	totalEpisodes := int(info["count"].(float64))
	originsByEpisode := make([]models.Episode, totalEpisodes)

	channel := make(chan models.Episode)

	for i := 1; i <= totalEpisodes; i++ {
		episodeID := strconv.Itoa(i)
		go episode.getOriginsByEachEpisode(baseURL, episodeID, channel)
	}

	for i := 0; i < totalEpisodes; i++ {
		episodeModel := <-channel
		originsByEpisode[episodeModel.ID-1] = episodeModel
	}

	return originsByEpisode
}

func (episode EpisodeAnalyzer) GetAllNames(channel chan string) {
	names := episode.analyzer.GetAllNames()
	for i := range names {
		channel <- names[i]
	}
	close(channel)
}

func (episode EpisodeAnalyzer) CountLetters(names []string, channel chan int) {
	count := episode.analyzer.CountLetters(names, "e")
	channel <- count
	close(channel)
}
