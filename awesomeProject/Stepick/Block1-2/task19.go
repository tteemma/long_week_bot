package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{},              // города с населением 10-99 тыс. человек
		100:  []string{"few", "reqd"}, // города с населением 100-999 тыс. человек
		1000: []string{},              // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"few":    1,
		"fewfew": 312,
	}

	for city := range cityPopulation {
		found := false
		for _, validCity := range groupCity[100] {
			if city == validCity {
				found = true
				break
			}
		}
		if !found {
			delete(cityPopulation, city)
			fmt.Println(city)
		}
	}
}
