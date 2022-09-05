package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var Input = []int{4, 5, 6, 7, 8, 9, 10}

func TestCheckIndex(t *testing.T) {

	a, b, c, d := CheckIndex(5, Input)
	log.Println("TestCheckIndex:")
	fmt.Println(a, b, c, d)
	if a != 5 || b != 2 || c != 6 || d != 3 {
		t.Fatalf("Expected output is inncorrect")
	}
}

func TestCheckIndexWithNoNextHighNum(t *testing.T) {

	a, b, c, d := CheckIndex(10, Input)
	log.Println("TestCheckIndexWithNoNextHighNum:")
	fmt.Println(a, b, c, d)

	if a != 10 || b != 7 || c != -1 || d != -1 {
		t.Fatalf("Expected output is inncorrect")
	}
}

func TestParseInputFile(t *testing.T) {
	//expectedOutput := make([]int{34, 100, 120, 50, 94, 22, 88, 1, 3, 4, 5, 6, 7, 8, 9, 0, 3456789, 35}, int)

	expectedOutput := [18]int{34, 100, 120, 50, 94, 22, 88, 1, 3, 4, 5, 6, 7, 8, 9, 0, 3456789, 35}
	a := ParseInputFile("./randomNumbers.txt")
	fmt.Println(expectedOutput)
	fmt.Println(a)
	if reflect.DeepEqual(expectedOutput, a) {

		t.Fatalf("File parsing is incorrect")
	}

}
