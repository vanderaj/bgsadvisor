package apihelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const systemEndPoint = "https://elitebgs.app/api/ebgs/v4/systems?id="

type systemResponse struct {
	Docs []struct {
		ID                      string    `json:"_id"`
		NameLower               string    `json:"name_lower"`
		X                       float64   `json:"x"`
		Y                       float64   `json:"y"`
		Z                       float64   `json:"z"`
		V                       int       `json:"__v"`
		Name                    string    `json:"name"`
		Government              string    `json:"government"`
		Allegiance              string    `json:"allegiance"`
		State                   string    `json:"state"`
		Security                string    `json:"security"`
		Population              int       `json:"population"`
		PrimaryEconomy          string    `json:"primary_economy"`
		ControllingMinorFaction string    `json:"controlling_minor_faction"`
		UpdatedAt               time.Time `json:"updated_at"`
		EddbID                  int       `json:"eddb_id"`
		Factions                []struct {
			Name      string `json:"name"`
			NameLower string `json:"name_lower"`
		} `json:"factions"`
		SecondaryEconomy string        `json:"secondary_economy"`
		SystemAddress    string        `json:"system_address"`
		Conflicts        []interface{} `json:"conflicts"`
	} `json:"docs"`
	Total int `json:"total"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

func getSystem(systemID string) systemResponse {
	resp, err := http.Get(factionEndpoint + systemID)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Todo struct
	var system systemResponse
	json.Unmarshal(bodyBytes, &system)

	return system
}
