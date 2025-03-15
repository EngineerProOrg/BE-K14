package fileservice

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"homework-14mar25/constants"
	"homework-14mar25/models"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]*models.Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		errorMessage := fmt.Sprintf("Cannot open file:%s", filename)
		return nil, errors.New(errorMessage)
	}
	defer file.Close()

	var persons []*models.Person
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 0 {
			errorMessage := fmt.Sprintf("File: %s incorrect format. File must contain '|'", filename)
			return nil, errors.New(errorMessage)
		}

		if len(parts[0]) == 0 || len(parts[1]) == 0 || len(parts[2]) == 0 {
			return nil, errors.New("invalid index")
		}

		personName := parts[0]
		personJobTitle := parts[1]
		personYearOfBirth, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			errorMessage := fmt.Sprintf("Cannot parse this value:%v into integer", parts[2])
			return nil, errors.New(errorMessage)
		}

		person := models.NewPerson(personName, personJobTitle, personYearOfBirth)
		persons = append(persons, person)
	}
	return persons, nil
}

func WriteFile(filename string, fileType constants.FileType, data interface{}) error {
	switch fileType {
	case constants.FileTxt:
		return writeTxt(filename, data)

	case constants.FileJson:
		return writeJson(filename, data)

	default:
		return errors.New("unsupported file type")
	}
}

func writeTxt(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch lines := data.(type) {
	case []string:
		for _, line := range lines {
			_, err := file.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	case []*models.Person:
		for _, person := range lines {
			line := fmt.Sprintf("%s|%s|%d", person.Name, person.JobTitle, person.YearOfBirth)
			_, err := file.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func writeJson(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// convert text into json format
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
