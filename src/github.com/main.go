package main

import (
	"fmt"
	"rick-and-morty-challenge/src/github.com/analyzers"
	"time"
)

func challenge1() {
	start := time.Now()

	listOfAnalyzers := make([]analyzers.Analyzer, 3)
	listOfAnalyzers[0] = &analyzers.LocationsAnalyzer{}
	listOfAnalyzers[1] = &analyzers.CharactersAnalyzer{}
	listOfAnalyzers[2] = &analyzers.EpisodeAnalyzer{}

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

	fmt.Println("Número de apariciones de la letra l en los nombres de todas las locaciones es: ")
	fmt.Println(<-listOfCountChannels[0])

	fmt.Println("Número de apariciones de la letra c en los nombres de todos los personajes es: ")
	fmt.Println(<-listOfCountChannels[1])

	fmt.Println("Número de apariciones de la letra e en los nombres de todos los episodios es: ")
	fmt.Println(<-listOfCountChannels[2])

	duration := time.Since(start)
	fmt.Println("Tiempo de ejecución: ")
	fmt.Println(duration)
}

func main() {
	challenge1()
}
