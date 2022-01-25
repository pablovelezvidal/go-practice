package main

import (
	"packages/countries"
)

func main() {
	countries.Add("Colombia")
	countries.Add("España")
	countries.Add("Italia")
	countries.Add("Thailandia")
	countries.Add("Francia")
	countries.Add("Brasil")
	countries.Add("Japon")

	err := countries.SetMyCountry("Brasil")

	if err != nil {
		panic(err)
	}

	countries.List()
}
