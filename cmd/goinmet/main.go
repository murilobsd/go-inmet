package main

import (
	"context"
	"fmt"

	"github.com/murilobsd/goinmet/pkg/inmet"
)

func main() {
	client := inmet.NewClient(nil)

	// list all stations
	stations, _, _ := client.Station.List(context.Background())
	for _, station := range stations {
		fmt.Printf("%v - %v", station.Code, station.Name)
	}
	fmt.Println("goinmet")
}
