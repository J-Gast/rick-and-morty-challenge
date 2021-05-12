package models

type Episode struct {
	ID              int
	EpisodeName     string
	OriginLocations map[string]bool
}

func (episode *Episode) Init() {
	episode.OriginLocations = make(map[string]bool)
}

func (episode Episode) GetListOriginLocations() []string {
	origins := make([]string, len(episode.OriginLocations))
	idx := 0
	for i := range episode.OriginLocations {
		origins[idx] = i
		idx++
	}
	return origins
}

func (episode Episode) GetNumOriginLocations() int {
	return len(episode.OriginLocations)
}
