package main

import (
	"os"
	"reflect"
	"rick-and-morty-challenge/src/github.com/analyzers"
	"rick-and-morty-challenge/src/github.com/models"
	"strconv"
	"testing"
)

var origins []models.Episode

func getOriginsByEpisodeMock() map[int]map[string]bool {
	originsByEpisode := make(map[int]map[string]bool)
	originsByEpisode[1] = map[string]bool{
		"Earth (C-137)":  true,
		"Bepis 9":        true,
		"Gromflom Prime": true,
		"unknown":        true,
		"Girvonesk":      true,
	}
	originsByEpisode[6] = map[string]bool{
		"Earth (C-137)":                 true,
		"unknown":                       true,
		"Earth (Replacement Dimension)": true,
		"Cronenberg Earth":              true,
	}
	originsByEpisode[11] = map[string]bool{
		"Earth (C-137)":                 true,
		"Earth (Replacement Dimension)": true,
		"unknown":                       true,
		"Girvonesk":                     true,
		"Testicle Monster Dimension":    true,
		"Bird World":                    true,
		"Gear World":                    true,
		"Fantasy World":                 true,
		"Planet Squanch":                true,
		"Bepis 9":                       true,
	}
	originsByEpisode[16] = map[string]bool{
		"Earth (C-137)":                 true,
		"Earth (Replacement Dimension)": true,
		"unknown":                       true,
		"Girvonesk":                     true,
		"Bird World":                    true,
		"Signus 5 Expanse":              true,
		"Alphabetrium":                  true,
		"Planet Squanch":                true,
		"Larva Alien's Planet":          true,
		"Árboles Mentirosos":            true,
	}
	originsByEpisode[21] = map[string]bool{
		"Earth (C-137)":                 true,
		"Earth (Replacement Dimension)": true,
		"unknown":                       true,
		"Bird World":                    true,
		"Glaagablaaga":                  true,
		"Planet Squanch":                true,
		"Larva Alien's Planet":          true,
		"Árboles Mentirosos":            true,
	}
	originsByEpisode[26] = map[string]bool{
		"Earth (C-137)":                 true,
		"Earth (Replacement Dimension)": true,
		"unknown":                       true,
		"Bird World":                    true,
		"Resort Planet":                 true,
	}
	originsByEpisode[31] = map[string]bool{
		"Earth (C-137)":                 true,
		"Earth (Replacement Dimension)": true,
		"unknown":                       true,
		"Mega Gargantuan Kingdom":       true,
	}
	originsByEpisode[37] = map[string]bool{
		"Earth (C-137)":                true,
		"Morty’s Story":                true,
		"Ricks’s Story":                true,
		"Story Train":                  true,
		"Tickets Please Guy Nightmare": true,
	}
	originsByEpisode[41] = map[string]bool{
		"Earth (C-137)":                 true,
		"Earth (Replacement Dimension)": true,
		"Bird World":                    true,
		"unknown":                       true,
		"Planet Squanch":                true,
		"Gromflom Prime":                true,
	}

	return originsByEpisode
}

func setup() {
	episodeAnalyzer := &analyzers.EpisodeAnalyzer{}
	origins = episodeAnalyzer.GetOrigins()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestNumOrigins(t *testing.T) {
	if len(origins) != 41 {
		t.Error("Num of episodes must be 41", 41, len(origins))
	}
}

func TestOrigins(t *testing.T) {
	originsByEpisode := getOriginsByEpisodeMock()
	for key, value := range originsByEpisode {
		if !reflect.DeepEqual(origins[key-1].OriginLocations, value) {
			t.Error("Origins are not equal in episode:"+strconv.Itoa(key), value, origins[key-1].OriginLocations)
		}
	}
}
