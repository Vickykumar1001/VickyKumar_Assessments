package main

import (
	"fmt"
	"strings"
)

type City struct {
	name        string
	temperature float64
	rainfall    float64
}

func main() {
	cities := []City{
		{"Delhi", 35.5, 150.0},
		{"Mumbai", 30.0, 220.5},
		{"Bengaluru", 25.0, 120.3},
		{"Kolkata", 32.0, 180.2},
		{"Chennai", 34.0, 200.1},
	}

	fmt.Println("Climate Data Analysis")
	fmt.Println("=======================")

	fmt.Println("1. City with highest and lowest temperatures")
	highest, lowest := findHighestAndLowestTemperature(cities)
	fmt.Printf("Highest Temperature: %s (%.1f°C)\n", highest.name, highest.temperature)
	fmt.Printf("Lowest Temperature: %s (%.1f°C)\n", lowest.name, lowest.temperature)

	fmt.Println("\n2. Average Rainfall")
	avgRainfall := calculateAverageRainfall(cities)
	fmt.Printf("Average Rainfall: %.2f mm\n", avgRainfall)

	fmt.Println("\n3. Filter Cities by Rainfall Threshold")
	var threshold float64
	fmt.Print("Enter rainfall threshold (mm): ")
	fmt.Scanln(&threshold)
	filterCitiesByRainfall(cities, threshold)

	fmt.Println("\n4. Search by City Name")
	fmt.Print("Enter city name to search: ")
	var searchName string
	fmt.Scanln(&searchName)
	searchCityByName(cities, strings.TrimSpace(searchName))
}
func findHighestAndLowestTemperature(cities []City) (City, City) {
	highest := cities[0]
	lowest := cities[0]

	for _, city := range cities {
		if city.temperature > highest.temperature {
			highest = city
		}
		if city.temperature < lowest.temperature {
			lowest = city
		}
	}
	return highest, lowest
}
func calculateAverageRainfall(cities []City) float64 {
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.rainfall
	}
	return totalRainfall / float64(len(cities))
}
func filterCitiesByRainfall(cities []City, threshold float64) {
	fmt.Printf("Cities with rainfall above %.1f mm:\n", threshold)
	for _, city := range cities {
		if city.rainfall > threshold {
			fmt.Printf("- %s: %.1f mm\n", city.name, city.rainfall)
		}
	}
}
func searchCityByName(cities []City, name string) {
	for _, city := range cities {
		if strings.EqualFold(city.name, name) {
			fmt.Printf("City: %s, Temperature: %.1f°C, Rainfall: %.1f mm\n", city.name, city.temperature, city.rainfall)
			return
		}
	}
	fmt.Println("City not found.")
}
