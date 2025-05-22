package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	ID   int
	Name string
	Job  string
	Dob  int
}

type UpdatePerson struct {
	Job string
}

func readPersonsFromFile(filename string) []Person {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var persons []Person
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 4 {
			fmt.Println("invalid line format:", line)
			continue
		}

		id, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		name := strings.TrimSpace(parts[1])
		job := strings.TrimSpace(parts[2])
		dob, err2 := strconv.Atoi(strings.TrimSpace(parts[3]))

		if err1 != nil || err2 != nil {
			fmt.Println("invalid data in line:", line)
			continue
		}

		person := Person{
			ID:   id,
			Name: name,
			Job:  job,
			Dob:  dob,
		}
		persons = append(persons, person)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v", err)
	}

	return persons
}

func insertPerson(persons *[]Person, person *Person) {
	newID := 0
	if len(*persons) > 0 {
		newID = (*persons)[len(*persons)-1].ID + 1
	}

	newPerson := Person{
		ID:   newID,
		Name: person.Name,
		Job:  person.Job,
		Dob:  person.Dob,
	}

	*persons = append(*persons, newPerson)

	fmt.Println("inserted person")
	fmt.Printf("ID: %d | Name: %s | Job: %s | Dob: %d\n", newPerson.ID, newPerson.Name, newPerson.Job, newPerson.Dob)
}

func updatePerson(persons *[]Person, id int, updatePerson *UpdatePerson) {
	var targetPerson *Person

	for i, person := range *persons {
		if person.ID == id {
			targetPerson = &(*persons)[i]
			break
		}
	}

	if targetPerson == nil {
		fmt.Println("Person not found")
		return
	}

	targetPerson.Job = updatePerson.Job

	fmt.Println("updated person")
	fmt.Printf("ID: %d | Name: %s | Job: %s | Dob: %d\n", targetPerson.ID, targetPerson.Name, targetPerson.Job, targetPerson.Dob)

}

func writePersonsToFile(persons *[]Person, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, person := range *persons {
		line := fmt.Sprintf("%d|%s|%s|%d\n", person.ID, person.Name, person.Job, person.Dob)
		writer.WriteString(line)
	}
	writer.Flush()
}

func printPersons(persons []Person) {
	for _, p := range persons {
		fmt.Printf("ID: %d | Name: %s | Job: %s | Dob: %d\n", p.ID, p.Name, p.Job, p.Dob)
	}
}

func main() {
	file := "./data.txt"
	persons := readPersonsFromFile(file)

	for {
		fmt.Println("=======================")
		fmt.Println("0. Exit")
		fmt.Println("1. Print list of persons")
		fmt.Println("2. Insert a person")
		fmt.Println("3. Update a person")
		fmt.Println("4. Write to file")

		fmt.Print("Input choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			printPersons(persons)
		case 2:
			insertPerson(&persons, &Person{
				Name: "DucTT",
				Job:  "Student",
				Dob:  2004,
			})
		case 3:
			updatePerson(&persons, 5, &UpdatePerson{
				Job: "SWE",
			})
		case 4:
			writePersonsToFile(&persons, file)
		case 0:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice!")
			continue
		}
	}
}
