package main

import (
	"fmt"
	"log"
)

func main() {
	parameters, err := NewParameters()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("parameters", parameters)
	}
}
