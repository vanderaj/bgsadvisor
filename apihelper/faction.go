package apihelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const factionEndpoint = "https://elitebgs.app/api/ebgs/v4/factions?id="
const opIdaFactionID = "5b6af33dd1b6a37c3c6ebddc"

type factionResponse struct {
	Docs []struct {
		ID              string    `json:"_id"`
		Name            string    `json:"name"`
		V               int       `json:"__v"`
		NameLower       string    `json:"name_lower"`
		UpdatedAt       time.Time `json:"updated_at"`
		Government      string    `json:"government"`
		Allegiance      string    `json:"allegiance"`
		FactionPresence []struct {
			SystemName      string  `json:"system_name"`
			SystemNameLower string  `json:"system_name_lower"`
			State           string  `json:"state"`
			Influence       float64 `json:"influence"`
			Happiness       string  `json:"happiness"`
			ActiveStates    []struct {
				State string `json:"state"`
			} `json:"active_states"`
			PendingStates    []interface{} `json:"pending_states"`
			RecoveringStates []struct {
				State string `json:"state"`
				Trend int    `json:"trend"`
			} `json:"recovering_states"`
			Conflicts []interface{} `json:"conflicts"`
			UpdatedAt time.Time     `json:"updated_at"`
		} `json:"faction_presence"`
		EddbID int `json:"eddb_id"`
	} `json:"docs"`
	Total int `json:"total"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

func getFaction() factionResponse {
	resp, err := http.Get(factionEndpoint + opIdaFactionID)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Todo struct
	var faction factionResponse
	json.Unmarshal(bodyBytes, &faction)

	return faction
}
