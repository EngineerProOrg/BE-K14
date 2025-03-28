package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	name string
	job  string
	dob  string
}

func main() {
	file := "text.txt"
	persons, err := readPersonsFromFile(file)
	terminateProgramIfErrorNotNil(err)
	persons = insertPerson(persons, "Huy", "SWE", "2000")
	printPersons(persons)

	err = updatePerson(persons, 0, "job", "Singer")
	terminateProgramIfErrorNotNil(err)
	printPersons(persons)

	err = writePersonsToFile(persons, file)
	terminateProgramIfErrorNotNil(err)
}
func terminateProgramIfErrorNotNil(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printPersons(persons []Person) {
	fmt.Println("Current persons list:")
	for i, person := range persons {
		fmt.Printf("[%d] %s\n", i, person)
	}
	fmt.Println()
}
func readPersonsFromFile(name string) ([]Person, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var persons []Person
	for scanner.Scan() {
		personFields := strings.Split(scanner.Text(), "|")
		if len(personFields) != 3 {
			return nil, fmt.Errorf("invalid record format: %s", scanner.Text())
		}
		persons = append(persons, Person{personFields[0], personFields[1], personFields[2]})
	}
	return persons, nil
}

func insertPerson(persons []Person, name, job, dob string) []Person {
	return append(persons, Person{name, job, dob})
}

func writePersonsToFile(persons []Person, file string) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, person := range persons {
		_, err := f.WriteString(fmt.Sprintf("%s|%s|%s\n", person.name, person.job, person.dob))
		if err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}
	return err
}
func updatePerson(persons []Person, index int, field, value string) (err error) {
	if index < 0 || index >= len(persons) {
		return fmt.Errorf("index out of range: %d", index)
	}
	switch field {
	case "name":
		persons[index].name = value
	case "job":
		persons[index].job = value
	case "dob":
		persons[index].dob = value
	default:
		return fmt.Errorf("invalid field: %s", field)
	}
	return err
}
