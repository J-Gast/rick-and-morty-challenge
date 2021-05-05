package analyzers

type LocationsAnalyzer struct {
	analyzer RickAnalyzer
	url      string
}

func (location *LocationsAnalyzer) Init() {
	location.url = "https://rickandmortyapi.com/api/location"
	location.analyzer = RickAnalyzer{URL: location.url}
}

func (location LocationsAnalyzer) GetAllNames(channel chan string) {
	names := location.analyzer.GetAllNames()

	for i := range names {
		channel <- names[i]
	}
	close(channel)
}

func (location LocationsAnalyzer) CountLetters(names []string) int {
	return location.analyzer.CountLetters(names, "l")
}
