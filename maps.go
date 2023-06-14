package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	/*create a map*/
	countryCapitalMap = make(map[string]string)

	/*insert key-value pairs in the map*/
	countryCapitalMap["Sri lanka"] = "Sri jayawardanapura kotte"
	countryCapitalMap["India"] = "New Dilhi"
	countryCapitalMap["America"] = "New york"
	countryCapitalMap["Austrailiya"] = "Sidny"

	/*print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("capital of country", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["India"]
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}
