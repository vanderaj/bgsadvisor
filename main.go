package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"

	apiclient "github.com/bgsadvisor/v2/client"
	"github.com/bgsadvisor/v2/client/operations"
	"github.com/bgsadvisor/v2/models"
	httptransport "github.com/go-openapi/runtime/client"
)

var faction string
var host string

var ticks models.TickTimesV5

func main() {

	fmt.Println("BGS Advisor v0.1")

	flag.StringVar(&faction, "faction", "Operation Ida", "Display BGS advice for this faction")
	flag.StringVar(&host, "host", "elitebgs.app", "Change the API host")
	flag.Parse()

	// create the transport
	transport := httptransport.New(host, "/api/ebgs/v5/", nil)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	apiclient.Default.SetTransport(transport)

	timeNow := time.Now().UnixNano() / 1e6
	// min time is 8 days ago, EliteBGS will only return up to 7
	timemin := strconv.FormatInt(timeNow, 10)
	timemax := strconv.FormatInt(timeNow, 10)

	// obtain tick data
	tick := operations.GetTicksParams{TimeMax: &timemax, TimeMin: &timemin}
	tick.SetTimeout(45 * time.Second)

	// make the request to get all items
	resp, err := client.Operations.GetTicks(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
	ticks = *resp.Payload[0]

	// obtain faction data

	// iterate and get system data

	// if tick is newer than system data, display a list of systems with outdated data, starting with controlled systems

	// Find the x (3 = default) systems with the most problems that need attention

	// display pending state changes that could need attention

	// display conflicts (system, asset, faction 1, faction 2, faction 1 days won, faction 2 days won)

	// display list of systems, inf, inf delta 1d, inf delta 7d, 1st faction, 2nd faction, last faction, order by margin of control

	// display bad state list

	// display min inf controllers
}
