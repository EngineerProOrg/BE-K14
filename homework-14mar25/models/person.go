package models

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name        string `json:"name"`
	JobTitle    string `json:"jobTitle"`
	YearOfBirth int64  `json:"yearOfBirth"`
}

func NewPerson(name string, jobTitle string, yearOfBirth int64) *Person {
	return &Person{
		Name:        name,
		JobTitle:    jobTitle,
		YearOfBirth: yearOfBirth,
	}
}

func PrintPersons(persons []*Person, message string) {
	fmt.Printf("\n **** List of Persons %s **** \n", message)
	for personIndex, person := range persons {
		fmt.Printf("%v Name: %s, Job: %s, Year: %d\n", personIndex+1, person.Name, person.JobTitle, person.YearOfBirth)
	}
}

func AddNewPerson(persons []*Person, name string, jobTitle string, yearOfBirth int64) []*Person {
	person := NewPerson(name, jobTitle, yearOfBirth)
	persons = append(persons, person)
	return persons
}

func AddNewPersonsFromKeyBoard(currentPersons []*Person) []*Person {
	var numberOfPersons int
	fmt.Println("Zero means skip method")
	fmt.Print("Enter your numbers of persons:")
	fmt.Scan(&numberOfPersons)

	if numberOfPersons == 0 {
		fmt.Println("Skip add new persons")
		return currentPersons
	}

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // remove '\n' after inputing number

	for index := range numberOfPersons {
		fmt.Printf("Person %v !", index+1)

		fmt.Print("\n Name:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Job title:")
		jobTitle, _ := reader.ReadString('\n')
		jobTitle = strings.TrimSpace(jobTitle)

		fmt.Print("Year of birth: ")
		yearStr, _ := reader.ReadString('\n')
		yearStr = strings.TrimSpace(yearStr)

		yearOfBirth, err := strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			fmt.Println("Invalid year format. Please enter a valid number.")
			index-- // go back and input valid data
			continue
		}
		person := NewPerson(name, jobTitle, int64(yearOfBirth))
		currentPersons = append(currentPersons, person)
	}
	return currentPersons
}

func UpdatePersonFromKeyBoard(updatedPersons []*Person) (*Person, error) {
	var personIndex int
	fmt.Println("Zero means skip method")
	fmt.Print("Enter your person index:")
	fmt.Scan(&personIndex)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // remove '\n' after inputing number

	personIndex = personIndex - 1

	if personIndex < 0 || personIndex >= len(updatedPersons) {
		errorMessage := fmt.Sprintf("Invalid number:%d", personIndex)
		return nil, errors.New(errorMessage)
	}

	fmt.Print("\nUpdated Name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Job title:")
	jobTitle, _ := reader.ReadString('\n')
	jobTitle = strings.TrimSpace(jobTitle)

	fmt.Print("Year of birth: ")
	yearStr, _ := reader.ReadString('\n')
	yearStr = strings.TrimSpace(yearStr)

	yearOfBirth, err := strconv.ParseInt(yearStr, 10, 64)

	if err != nil {
		errorMessage := fmt.Sprintf("Invalid year of birth:%d", yearOfBirth)
		return nil, errors.New(errorMessage)
	}

	updatedPersons[personIndex].Name = name
	updatedPersons[personIndex].JobTitle = jobTitle
	updatedPersons[personIndex].YearOfBirth = yearOfBirth

	return updatedPersons[personIndex], nil
}
