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
func checkNumInFile(inputNum int, filename string) (int, int) {

	fmt.Println("Processing the file.....")

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

	//convert string slice into int slice
	IntSlice := make([]int, len(strArr))
	for i, stt := range strArr {
		IntSlice[i], _ = strconv.Atoi(stt)
	}

	//sort map by key
	m, _ := make(map[int]int), err
	for i := 0; i < len(IntSlice); i++ {
		m[IntSlice[i]] = i + 1
	}
	fmt.Println("Position of next highest elements:", m[inputNum+1])
	//fmt.Println(inputNum)

	n, ok := m[inputNum]
	if ok == true {
		return n, m[inputNum+1]
	}
	//return -1 if the input is not present
	return -1, 0
}

func main() {

	var inputNum int
	fmt.Println("Enter number you want to find:")
	if _, err := fmt.Scan(&inputNum); err != nil {
		log.Printf("Enter valid number ") //enter cases if number entered is not string.
		return
	}

	filename := "./randomNumbers.txt"
	result, nextVal := checkNumInFile(inputNum, filename)

	if result == -1 || result == 0 {
		fmt.Println("The number is not present.")
	} else {
		fmt.Println("Position of number:", result)
		fmt.Println("Position of next highest number:", nextVal)
	}

}
