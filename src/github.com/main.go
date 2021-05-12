package main

import (
	"bufio"
	"fmt"
	"os"
	"rick-and-morty-challenge/src/github.com/analyzers"
	"rick-and-morty-challenge/src/github.com/utils"
	"strconv"
	"time"
)

func Challenge1() {
	start := time.Now()
	analyzer := &analyzers.RickAnalyzer{}
	listOfCountChannels := analyzer.Analyze()
	fmt.Println("Número de apariciones de la letra l en los nombres de todas las locaciones es: ")
	fmt.Println(listOfCountChannels[0])

	fmt.Println("Número de apariciones de la letra c en los nombres de todos los personajes es: ")
	fmt.Println(listOfCountChannels[1])

	fmt.Println("Número de apariciones de la letra e en los nombres de todos los episodios es: ")
	fmt.Println(listOfCountChannels[2])

	duration := time.Since(start)
	fmt.Print("Tiempo de ejecución: ")
	fmt.Println(duration)
}

func Challenge2() {
	start := time.Now()
	episodeAnalyzer := &analyzers.EpisodeAnalyzer{}
	origins := episodeAnalyzer.GetOrigins()
	for i := 0; i < len(origins); i++ {
		episodeModel := origins[i]
		origins := episodeModel.GetListOriginLocations()
		fmt.Println("El episodio " + strconv.Itoa(episodeModel.ID) + " cuenta con " + strconv.Itoa(episodeModel.GetNumOriginLocations()) + " lugares de origen de los personajes")
		fmt.Println(origins)
		fmt.Println("--------------------------")
	}
	duration := time.Since(start)
	fmt.Print("Tiempo de ejecución: ")
	fmt.Println(duration)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cleaner := utils.Cleaner{}
	cleaner.Init()
	for {
		cleaner.CallClear()
		fmt.Println("---------------------")
		fmt.Println("¡Bienvenido al Rick and Morty Challenge!")
		fmt.Println("Presiona 1 para ejecutar el Challenge 1.")
		fmt.Println("Presiona 2 para ejecutar el Challenge 2.")
		fmt.Println("---------------------")
		opt, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
		switch opt {
		case '1':
			Challenge1()
			fmt.Print("Presiona Enter para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		case '2':
			Challenge2()
			fmt.Print("Presiona Enter para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		}
	}
}
