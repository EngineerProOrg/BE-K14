package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name        string
	Occupation  string
	YearOfBirth int
}

const (
	NAME          = "NAME"
	OCCUPATION    = "OCCUPATION"
	YEAR_OF_BIRTH = "YEAR_OF_BIRTH"
)

func main() {
	people := readDataFromFile("data.txt")
	fmt.Println("Initial data")
	fmt.Println(people)

	people = insertPerson("Chris", "Engineer", 1990, people)
	fmt.Println("After inserting Chris")
	fmt.Println(people)

	fmt.Println("Before updating the first person job")
	fmt.Println(people)
	people[0] = updatePerson(0, OCCUPATION, "Vibe coder", people)
	fmt.Println("After updating the first person job")
	fmt.Println(people)
	fmt.Println("Updating the second person year of birth")
	people[1] = updatePerson(1, YEAR_OF_BIRTH, 1983, people)
	fmt.Println(people)

    fmt.Println("Overwriting the file with updated data")
    overwriteFile("data.txt", people)
}

func readDataFromFile(filename string) []Person {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	people := []Person{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		data := strings.Split(text, "|")
		yearOfBirth, _ := strconv.Atoi(data[2])
		people = append(people, Person{
			Name:        data[0],
			Occupation:  data[1],
			YearOfBirth: yearOfBirth,
		})
	}
	return people
}

func insertPerson(name, job string, yearOfBirth int, people []Person) []Person {
	person := Person{
		Name:        name,
		Occupation:  job,
		YearOfBirth: yearOfBirth,
	}
	people = append(people, person)
	return people
}

func updatePerson(index int, field string, value interface{}, people []Person) Person {
	if index < 0 || index >= len(people) {
		fmt.Println("Invalid index")
	}
	person := people[index]
	switch field {
	case NAME:
		name, ok := value.(string)
		if ok {
			person.Name = name
		} else {
			fmt.Println("Name has to be a string")
		}
	case OCCUPATION:
		occupation, ok := value.(string)
		if ok {
			person.Occupation = occupation
		} else {
			fmt.Println("Occupation has to be a string")
		}
	case YEAR_OF_BIRTH:
		yob, ok := value.(int)
		if ok {
			if yob < 0 {
				fmt.Println("Year of birth has to be a positive integer")
			} else {
				person.YearOfBirth = yob
			}
		} else {
			fmt.Println("Year of birth has to be an integer")
		}
	default:
		fmt.Println("Invalid field")
	}
	return person
}

func overwriteFile(filename string, people []Person) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, person := range people {
		writeData := fmt.Sprintf("%s|%s|%d\n", person.Name, person.Occupation, person.YearOfBirth)
		_, err := file.Write([]byte(writeData))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Overwriting the file with data")
}
