package apihelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const tickEndPoint = "https://elitebgs.app/api/ebgs/v4/ticks"

// TickResponse contains the last tick
type TickResponse struct {
	ID          string `json:"_id"`
	TickTime    string `json:"time"`
	UpdatedTime string `json:"updated_at"`
	Unused      int    `json:"__v"`
}

func getTick() TickResponse {
	resp, err := http.Get(tickEndPoint)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Todo struct
	var tick TickResponse
	json.Unmarshal(bodyBytes, &tick)

	return tick
}
