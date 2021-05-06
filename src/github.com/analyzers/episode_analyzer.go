package analyzers

type EpisodeAnalyzer struct {
	analyzer RickAnalyzer
	url      string
}

func (episode *EpisodeAnalyzer) Init() {
	episode.url = "https://rickandmortyapi.com/api/episode"
	episode.analyzer = RickAnalyzer{URL: episode.url}
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
