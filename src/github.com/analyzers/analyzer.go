package analyzers

type Analyzer interface {
	Init()
	GetAllNames(channel chan string)
	CountLetters(names []string, channel chan int)
}
