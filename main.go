package main

import (
	"Sierrachart2Tradervue/dtc"
	"Sierrachart2Tradervue/dtc2tv"
	"Sierrachart2Tradervue/tradervue"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {

	// required arguments
	username := flag.String("tv_username", "", "Tradervue username")
	password := flag.String("tv_password", "", "Tradervue password")
	tradeAccount := flag.String("sc_account", "", "Sierrachart trade account")
	numberOfDays := flag.Int("sc_days", 0, "Number of days to request fills for")

	// optional arguments
	tagsRaw := flag.String("tv_tags", "", "Comma separated list of tags to apply to all executions")
	accountTag := flag.String("tv_account_tag", "", "Account tag to apply to all executions")

	// make sure required arguments are specified
	flag.Parse()
	if *username == "" || *password == "" || *tradeAccount == "" || *numberOfDays == 0 {
		panic("Missing argument(s). Please specify username, password, account and days.")
	}

	// create SierraChart DTC client and logon to the local DTC server
	c, err := dtc.NewClient()
	if err != nil {
		panic(err)
	}

	// request historical fills from SierraChart
	log.Printf("Requesting %d day(s) of historical fills for trade account `%s` from SierraChart\n", *numberOfDays, *tradeAccount)
	fills, err := c.RequestHistoricalFills(*tradeAccount, *numberOfDays)
	if err != nil {
		panic(err)
	}
	log.Printf("Received %d fills from Sierrachart\n", len(fills))

	//convert SierraChart fills to Tradervue executions
	var executions []tradervue.Execution
	for _, fill := range fills {
		tve, err := dtc2tv.ExecutionFromDTCOrderFill(fill)
		if err != nil {
			panic(fmt.Sprintf("Error during conversion %s", err.Error()))
		}
		executions = append(executions, tve)
	}

	//import into TV
	log.Printf("Importing into %d executions into Tradervue...\n", len(executions))
	tags := strings.Split(*tagsRaw, ",")
	err = tradervue.Import(executions, *username, *password, tags, accountTag)
	if err != nil {
		panic(fmt.Sprintf("Tradervue import failed with %s", err.Error()))
	}
}
