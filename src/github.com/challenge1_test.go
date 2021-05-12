package main

import (
	"rick-and-morty-challenge/src/github.com/analyzers"
	"testing"
)

func TestAnalyzers(t *testing.T) {
	analyzer := &analyzers.RickAnalyzer{}
	counters := analyzer.Analyze()
	if counters[0] != 73 {
		t.Error("La letra L en los nombres de todas las locaciones debe ser 73", 73, counters[0])
	}
	if counters[1] != 394 {
		t.Error("La letra C en los nombres de todos los personajes debe ser 394", 394, counters[1])
	}
	if counters[2] != 71 {
		t.Error("La letra E en los nombres de todos los episodios debe ser 71", 71, counters[2])
	}
}
