package main

import (
	"log"

	"github.com/koron/go-ssdp"
)

func searchForRokus() chan []ssdp.Service {
	resultChan := make(chan []ssdp.Service)

	go func() {
		defer close(resultChan)

		discoveredRokus, err := ssdp.Search("roku:ecp", 1, "")
		if err != nil {
			log.Println("Error discovering Rokus:", err)
			return
		}

		if discoveredRokus != nil {
			resultChan <- discoveredRokus // Send results back to the main thread
			// Deduplicate based on LOCATION
			// Needed because the SSDP code runs on EVERY interface :)
			// So, if you have WiFi and Ethernet enabled, it will show two callbacks from your Roku TV.
			// The code is just that *good*
			locationMap := make(map[string]ssdp.Service)
			for _, roku := range discoveredRokus {
				locationMap[roku.Location] = roku
			}

			// Convert map back to a slice
			uniqueRokus := make([]ssdp.Service, 0, len(locationMap))
			for _, roku := range locationMap {
				uniqueRokus = append(uniqueRokus, roku)
			}
			resultChan <- uniqueRokus // Send results back to the main thread
		} else {
			resultChan <- nil // No Rokus found, send nil
		}
	}()

	return resultChan
}
