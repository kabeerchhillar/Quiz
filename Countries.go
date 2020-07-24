package main

import (
	"CountryAbb"
	"CountryCapital"
	"CountryLanguage"
	"CountryReligion"
	"fmt"
)

func main() {

	fmt.Println("What do you want to play \n" +
		"type 1 to play Country Capital \n" +
		"type 2 to play Country Language \n" +
		"type 3 to play Country Abbreviation\n" +
		"type 4 to play Country Religion")

	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {

		var countryName string
		fmt.Scanln(&countryName)

		capital := CountryCapital.GetTheCapital(countryName,CountryCapital.CapitalReader())
		fmt.Println(capital)

	}
	if choice == 2 {

		if CountryLanguage.CheckYourLanguageKnowledge(CountryLanguage.LanguageReader()) == false {
			fmt.Println("You are gone ... you got so may tries")
		}

	}
	if choice == 3 {

		if CountryAbb.CheckYourAbbbreviationLanguage(CountryAbb.AbbreviationReader()) == false {
			fmt.Println("You are gone ... you got so may tries")
		}

	}
	if choice == 4 {

		if CountryReligion.CheckYourReligionLanguage(CountryReligion.ReligionReader()) == false {
			fmt.Println("You are gone ... you got so may tries")
		}

	}

}