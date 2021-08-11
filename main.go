package main

import (
	"fmt"
	"os"

	"main.go/Common"
	"main.go/TrackingService"
)

func main() {
	parameters, err := Common.NewParameters()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		trackService := TrackingService.NewTrackingService(parameters)
		err = trackService.Track()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
