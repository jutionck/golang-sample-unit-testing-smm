package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	myCalc := Calculator{Num1: 4, Num2: 4}
	actual := myCalc.Add()
	expected := 9
	assert.Equal(t, expected, actual, "it should return 8")

}
func TestCalculator_Sub(t *testing.T) {
	myCalc := Calculator{Num1: 5, Num2: 4}
	actual := myCalc.Sub()
	expected := 1
	assert.Equal(t, expected, actual, "it should return 1")
}

func TestMainApp(t *testing.T) {
	os.Args = []string{"-num1", "2", "-num2", "3", "-opr", "add"}
	main()
}

//func TestMainApp_Sub(t *testing.T) {
//	//os.Args[5] = "sub"
//	os.Args = []string{"-num1", "1", "-num2", "3", "-opr", "sub"}
//	main()
//}
