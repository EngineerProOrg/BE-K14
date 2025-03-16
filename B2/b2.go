package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Job  string
	Dob  uint16
}

func getItemAtIndexOrDefault[T any](slice []T, index int, defaultValue T) T {
	if index >= 0 && index < len(slice) {
		return slice[index]
	}
	return defaultValue
}

func main() {
	var people []Person = make([]Person, 0)
	var filePath string = "./text.txt"
	people = readFromFile(filePath)
	fmt.Println("People: ", people)

	//Insert dữ liệu vào slice. Input: name,job,dob. Output: thêm dữ liệu vào cuối slice
	people = insertPersonAtLast(people, Person{Name: "An", Job: "Plumber", Dob: 1987})

	//Update dữ liệu trong slice: Input: index, field (name/job/dob), value. Output: update field=value ch
	updatePersonAtIndex(people, 1, "Name", "AN")
	fmt.Println("People: ", people)
	//updatePersonAtIndex(people, 1, "Name2", "AN")

	//Ghi đè dữ liệu từ slice vào file.
	err := writeToFile(filePath, people)
	if err != nil {
		panic(err)
	}
}

func readFromFile(filePath string) []Person {
	var people []Person = make([]Person, 0)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), "|")
		name, job, dob := getItemAtIndexOrDefault(x, 0, ""), getItemAtIndexOrDefault(x, 1, ""), getItemAtIndexOrDefault(x, 2, "")
		i, err := strconv.ParseUint(dob, 10, 16)
		if err != nil {
			fmt.Println("Error", err)
			return nil
		}
		people = append(people, Person{Name: name, Job: job, Dob: uint16(i)})
	}
	return people
}

func writeToFile(fileName string, slice []Person) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // Ensure all buffered writes are flushed before closing

	for _, item := range slice {
		str := fmt.Sprintf("%s|%s|%d", item.Name, item.Job, item.Dob)
		_, err := writer.WriteString(str + "\n") // Append newline character
		if err != nil {
			return err
		}
	}
	return nil
}

func insertPersonAtLast(people []Person, p Person) []Person {
	return append(people, p)
}
func updatePersonAtIndex[T any](people []Person, index int, fieldName string, value T) {
	if len(people) <= index {
		panic("Index out of range")
	}
	person := &people[index]
	val := reflect.ValueOf(person).Elem()
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		panic(fmt.Sprintf("Field %s is not found", fieldName))
	}
	newValue := reflect.ValueOf(value)
	if newValue.Type().ConvertibleTo(field.Type()) {
		field.Set(newValue.Convert(field.Type()))
	} else {
		panic(fmt.Sprintf("Cannot assign value of type %T to field %s", value, fieldName))
	}
}
