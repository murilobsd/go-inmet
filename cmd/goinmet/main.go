package main

import (
	"context"
	"fmt"

	"github.com/murilobsd/goinmet/pkg/inmet"
)

func main() {
	client := inmet.NewClient(nil)

	// list all stations
	// stations, _, _ := client.Station.List(context.Background())
	// for _, station := range stations {
	// 	fmt.Printf("%v - %v\n", station.Code, station.Name)
	// }

	// list all culture cycle
	// culturesCycle, _, _ := client.Culture.CycleList(context.Background())
	// for _, culture := range culturesCycle {
	// 	fmt.Printf("%v\n", culture.Description)
	// }

	// list all soils type
	// soilsType, _, _ := client.Soil.TypeList(context.Background())
	// for _, soil := range soilsType {
	// 	fmt.Printf("%v\n", soil.Description)
	// }

	// get cad final
	cadFinal, _, _ := client.CAD.GetFinal(context.Background(), 1, "2036")
	fmt.Println("Cad Final: ", *cadFinal)
}
