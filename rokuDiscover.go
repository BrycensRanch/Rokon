package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"

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
		if discoveredRokus == nil {
			return
		}

		// Deduplicate based on LOCATION
		// Needed because the SSDP code runs on EVERY interface :)
		// So, if you have WiFi and Ethernet enabled, it will show two callbacks from your Roku TV.
		// The code is just that *good*
		rokuMap := make(map[string]ssdp.Service)
		for _, roku := range discoveredRokus {
			rokuIP, err := getHostFromLocation(roku.Location)
			if err != nil {
				log.Panicf("Failed to get host from roku location. %s", roku.Location)
			}
			rokuMap[rokuIP] = roku
		}

		uniqueRokus := make([]ssdp.Service, 0, len(rokuMap))
		for _, roku := range rokuMap {
			uniqueRokus = append(uniqueRokus, roku)
		}
		resultChan <- uniqueRokus
	}()

	return resultChan
}

func logNetworkInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error fetching network interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		status := "down"
		if iface.Flags&net.FlagUp != 0 {
			status = "up"
		}

		var ifaceType string
		switch {
		case iface.Flags&net.FlagLoopback != 0:
			ifaceType = "loopback"
		case strings.Contains(iface.Name, "en") || strings.Contains(iface.Name, "eth"):
			ifaceType = "Ethernet"
		case strings.Contains(iface.Name, "wl"):
			ifaceType = "Wi-Fi"
		default:
			ifaceType = "Unknown"
		}

		log.Printf("Interface: %s, Status: %s, Type: %s\n", iface.Name, status, ifaceType)
	}
}

func getHostFromLocation(location string) (string, error) {
	parsedURL, err := url.Parse(location)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	host := parsedURL.Hostname()

	return host, nil
}
