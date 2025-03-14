package main

import (
	"fmt"

	"homework-14mar25/constants"
	"homework-14mar25/models"
	"homework-14mar25/services/fileservice"
)

func main() {
	persons, err := fileservice.ReadLines(constants.FileTxtPersons)
	if err != nil {
		fmt.Print(err)
		return
	}

	models.PrintPersons(persons, "after reading file")

	persons = models.AddNewPerson(persons, "Khanh Vuong Bui", "Senior .NET Developer", 1992)
	models.PrintPersons(persons, "after adding new person")

	persons = models.AddNewPersonsFromKeyBoard(persons)
	models.PrintPersons(persons, "after adding persons from keyboard")

	models.UpdatePersonFromKeyBoard(persons)
	models.PrintPersons(persons, "after updating person from keyboard")

	err = fileservice.WriteFile(constants.FileJsonPersons, constants.FileJson, persons)
	if err != nil {
		fmt.Print(err)
		return
	}

	outputPerson := fmt.Sprintf("output_%s", constants.FileTxtPersons)
	err = fileservice.WriteFile(outputPerson, constants.FileTxt, persons)
	if err != nil {
		fmt.Print(err)
		return
	}
}
