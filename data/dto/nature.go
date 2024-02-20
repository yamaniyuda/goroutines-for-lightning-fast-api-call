package dto

type NatureResult struct {
	Name       string       `json:"name"`
	StatChange []StatChange `json:"pokeathlon_stat_changes"`
}

type StatChange struct {
	MaxChange float32        `json:"max_change"`
	stats     PokeathlonStat `json:pokeathlon_stat`
}

type PokeathlonStat struct {
	Name string `json:"name"`
}
