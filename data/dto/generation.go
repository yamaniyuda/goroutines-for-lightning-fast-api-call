package dto

type GenerationResult struct {
	Name          string          `json:"name"`
	VersionGroups []VersionGroups `json:"version_groups"`
}

type VersionGroups struct {
	Name string `json:"name"`
	Url  string `json:"url`
}
