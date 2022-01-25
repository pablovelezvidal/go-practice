package countries

import (
	"errors"
	"fmt"
)

var collection []string
var myCountry string
var errorNotFound error = errors.New("country not found")

// Add agrega el valor de pais a la collection
func Add(country string) {
	collection = append(collection, country)
}

// SetMyCountry selecciona un pais como mi pais de residencia
func SetMyCountry(country string) error {
	if !isInCollection(country) {
		return errorNotFound
	}
	for _, c := range collection {
		if c == country {
			myCountry = country
		}
	}
	return nil
}

func List() {
	for i, country := range collection {
		myCountryMsg := ""
		if country == myCountry {
			myCountryMsg = "[My Country]"
		}
		fmt.Printf("%d - %s %s \n", i+1, country, myCountryMsg)
	}
}

func isInCollection(country string) bool {
	for _, c := range collection {
		if c == country {
			return true
		}
	}
	return false
}
