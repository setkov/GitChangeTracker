package main

import (
	"fmt"
	"log"

	"main.go/Common"
	"main.go/TrackingService"
)

func main() {
	parameters, err := Common.NewParameters()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("parameters", parameters)
	}

	trackService := TrackingService.NewTrackingService(parameters)
	err = trackService.Track()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("track commit success")
	}

}
