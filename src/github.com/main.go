package main

import (
	"fmt"
	"rick-and-morty-challenge/src/github.com/analyzers"
	"time"
)

func main() {
	start := time.Now()
	locationsAnalyzer := analyzers.LocationsAnalyzer{}
	charactersAnalyzer := analyzers.CharactersAnalyzer{}
	episodesAnalyzer := analyzers.EpisodeAnalyzer{}
	locationsAnalyzer.Init()
	charactersAnalyzer.Init()
	episodesAnalyzer.Init()

	var locationsNames []string
	var charactersNames []string
	var episodeNames []string
	locationsChan := make(chan string)
	charactersChan := make(chan string)
	episodesChan := make(chan string)

	go locationsAnalyzer.GetAllNames(locationsChan)
	go charactersAnalyzer.GetAllNames(charactersChan)
	go episodesAnalyzer.GetAllNames(episodesChan)

	for name := range locationsChan {
		locationsNames = append(locationsNames, name)
	}
	for name := range charactersChan {
		charactersNames = append(charactersNames, name)
	}
	for name := range episodesChan {
		episodeNames = append(episodeNames, name)
	}

	numL := locationsAnalyzer.CountLetters(locationsNames)
	numC := charactersAnalyzer.CountLetters(charactersNames)
	numE := episodesAnalyzer.CountLetters(episodeNames)

	fmt.Println("Número de apariciones de la letra l en los nombres de todas las locaciones es: ")
	fmt.Println(numL)

	fmt.Println("Número de apariciones de la letra c en los nombres de todos los personajes es: ")
	fmt.Println(numC)

	fmt.Println("Número de apariciones de la letra e en los nombres de todos los episodios es: ")
	fmt.Println(numE)

	duration := time.Since(start)
	fmt.Println("Tiempo de ejecución: ")
	fmt.Println(duration)
}
