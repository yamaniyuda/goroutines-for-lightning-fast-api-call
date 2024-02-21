package repository

import (
	"encoding/json"
	"fmt"
	"goroutines-for-lightning-fast-api-call/data/dto"
	"io"
	"net/http"
)

const NATURE_URL = "https://pokeapi.co/api/v2/nature/calm\""

func GetNatureData() {
	natureReq, _ := http.NewRequest(http.MethodGet, NATURE_URL, nil)
	natureRes, _ := http.DefaultClient.Do(natureReq)
	defer natureRes.Body.Close()

	body, _ := io.ReadAll(natureRes.Body)
	var responseObject dto.NatureResult
	json.Unmarshal(body, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.StatChange))

	for _, statChange := range responseObject.StatChange {
		fmt.Println(statChange.Stats.Name)
	}
}
