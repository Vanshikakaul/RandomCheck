package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// checkNumInFile function checks whether the input number
// is present in the file or not.

func ParseInputFile(filename string) ([]int, error) {
	fmt.Println("Processing the file.....")
	var (
		data []byte
		err  error
	)
	// read the file exist the fn if the file is not present.
	if data, err = ioutil.ReadFile(filename); err != nil {
		log.Fatal(err)
	}

	//convert the byte data to string
	s := string(data)
	fmt.Println("s", s)

	//split the string using the ',' delimiter
	//into string array
	strArr := strings.Split(s, ",")

	//convert string slice into int slice
	inputArr := make([]int, len(strArr))
	for i, stt := range strArr {
		inputArr[i], err = strconv.Atoi(stt)
		if err != nil {
			return inputArr, err
		}

	}
	return inputArr, nil
}

// checking index of input number and next highest number
func CheckIndex(inputNum int, inputArr []int) (int, int, int, int) {

	var (
		nextHighNum      int
		indexNextHighNum int
		isfoundgdnum     bool
	)

	// define map
	m := make(map[int]int)
	//coverting array to map where key = number, value= position
	for i := 0; i < len(inputArr); i++ {
		m[inputArr[i]] = i + 1

		//take the greater num of the array to find the next greater number.
		if nextHighNum < inputArr[i] {
			nextHighNum = inputArr[i]
		}
	}

	//Get position of the Number from the sorted map
	//and check if number exist or not

	if inputNumIndex, ok := m[inputNum]; ok {
		fmt.Println(fmt.Sprintf("Index of the input num: %+v is at position: %+v", inputNum, inputNumIndex))
		for key, value := range m {

			if key > inputNum && key <= nextHighNum {
				nextHighNum = key
				indexNextHighNum = value
				isfoundgdnum = true
			}
		}
		if isfoundgdnum {
			fmt.Println(fmt.Sprintf("Index of the next highest num: %+v is at position: %+v", nextHighNum, indexNextHighNum))
			return inputNum, inputNumIndex, nextHighNum, indexNextHighNum
		} else {
			fmt.Println("No next highest number for input value")
			return inputNum, inputNumIndex, -1, -1
		}

	} else {
		fmt.Println("Input number not found")
		return inputNum, -1, -1, -1
	}

}

func main() {

	var inputNum int
	fmt.Println("Enter number you want to find:")
	if _, err := fmt.Scan(&inputNum); err != nil {
		log.Printf("Please enter valid number ")
		return
	}

	filename := "./randomNumbers.txt"
	fileArr, err := ParseInputFile(filename)
	if err != nil {
		log.Println("The file contain Invalid data")
		return
	}

	CheckIndex(inputNum, fileArr)

}
