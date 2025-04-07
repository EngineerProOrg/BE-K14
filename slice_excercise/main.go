package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Tạo struct Person gồm các field tương ứng (tạm thời cứ để hết là string cho dễ nhé, bạn nào làm đc int thì tốt).
	Viết chương trình go bao gồm các chức năng:
	- Đọc dữ liệu từ file và lưu vào 1 slice Person
	- Insert dữ liệu vào slice. Input: name,job,dob. Output: thêm dữ liệu vào cuối slice
	- Update dữ liệu trong slice: Input: index, field (name/job/dob), value. Output: update field=value cho phần tử index.
	- Ghi đè dữ liệu từ slice vào file.

	data:
		Tom|Software engineer|1995
		John Snow|Teacher|1997
		Maria Onitsuka|Actor|1993
		Emil|Football player|1987
*/

type Person struct {
	name string
	job  string
	dob  int
}

func setPerson(name string, job string, dob int) *Person {
	return &Person{name, job, dob}
}

func readData(nameOfFile string) (result []Person, err error) {
	file, err := os.Open(nameOfFile)
	if err != nil {
		fmt.Println("Error when opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//todo: handle data from file
	for scanner.Scan() {
		dataLine := scanner.Text()
		fields := strings.Split(dataLine, "|")
		if len(fields) != 3 {
			return nil, errors.New("invalid data")
		}
		name := strings.TrimSpace(fields[0])
		job := strings.TrimSpace(fields[1])
		dob, err := strconv.Atoi(strings.TrimSpace(fields[2]))
		if err != nil {
			return nil, errors.New("invalid dob data")
		}
		person := setPerson(name, job, dob)
		result = append(result, *person)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error when scanning file:", err)
	}
	return result, nil
}

func updateData(result *[]Person, index int, person *Person) {
	(*result)[index] = *person
}

func overrideData(nameOfFile string, persons []Person) {
	var updateDatas string
	for _, person := range persons {
		updateDatas += person.name + "|" + person.job + "|" + strconv.Itoa(person.dob) + "\n"
	}
	err := os.WriteFile(nameOfFile, []byte(updateDatas), 0644)
	if err != nil {
		fmt.Println("Error re-writing file:", err)
	}
}

func main() {
	nameOfFile := os.Args[1]
	result, err := readData(nameOfFile)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Data before: ", result)
	newData := Person{
		name: "Taylor",
		job:  "Singer",
		dob:  1999,
	}
	updateData(&result, 1, &newData)
	fmt.Println("Data after: ", result)
	overrideData(nameOfFile, result)
}
