package repository

import (
	"encoding/json"
	"fmt"
	"goroutines-for-lightning-fast-api-call/data/dto"
	"io"
	"net/http"
)

const GENERATION_URL = "https://pokeapi.co/api/v2/generation/generation-ii"

func GetGenerationData() {
	generationReq, _ := http.NewRequest(http.MethodGet, GENERATION_URL, nil)
	generationRes, _ := http.DefaultClient.Do(generationReq)
	defer generationRes.Body.Close()

	body, _ := io.ReadAll(generationRes.Body)
	var responseObject dto.GenerationResult
	json.Unmarshal(body, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.VersionGroups))

	for _, versionGroup := range responseObject.VersionGroups {
		fmt.Println(versionGroup.Name)
	}
}
