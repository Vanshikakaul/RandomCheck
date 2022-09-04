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
func checkNumInFile(inputNum int, filename string) int {

	fmt.Println("Processing the file.....")

	inputStr := strconv.Itoa(inputNum)

	var data []byte
	var err error
	// read the file exist the fn if the file is not present.
	if data, err = ioutil.ReadFile(filename); err != nil {
		log.Fatal(err)
	}

	//convert the byte data to string
	s := string(data)

	//split the string using the ',' delimiter
	//into string array
	strArr := strings.Split(s, ",")

	// loop over the array and check whether the
	//input is present.
	for i := 0; i < len(strArr); i++ {

		if strArr[i] == inputStr {
			return i
		}

	}

	//return -1 if the input is not present
	return -1

}

func main() {

	var inputNum int

	fmt.Println("Enter number you want to find:")
	fmt.Scanf("%d\n", &inputNum)

	filename := "./randomNumbers.txt"

	result := checkNumInFile(inputNum, filename)

	if result == -1 {
		fmt.Println("The number is not present.")
	} else {
		fmt.Println("The number is present at the postion", result+1)
	}

}
