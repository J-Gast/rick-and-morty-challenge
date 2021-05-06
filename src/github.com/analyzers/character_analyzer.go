package analyzers

type CharactersAnalyzer struct {
	analyzer RickAnalyzer
	url      string
}

func (character *CharactersAnalyzer) Init() {
	character.url = "https://rickandmortyapi.com/api/character"
	character.analyzer = RickAnalyzer{URL: character.url}
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
