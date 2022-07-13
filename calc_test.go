package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CalculatorTestSuite struct {
	suite.Suite
	calcService Calculator
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calcService = Calculator{}
}

func (suite *CalculatorTestSuite) TestCalculator_Add() {
	testTable := []struct {
		num1 int
		num2 int
		res  int
	}{
		{4, 4, 8},
		{-4, 5, -1},
	}
	for _, table := range testTable {
		suite.calcService.Num1 = table.num1
		suite.calcService.Num2 = table.num2
		actual, err := suite.calcService.Add()
		if err != nil {
			assert.NotNil(suite.T(), err)
			assert.Equal(suite.T(), table.res, actual)
		} else {
			assert.Nil(suite.T(), err)
			assert.Equal(suite.T(), table.res, actual)

		}
	}
}
func (suite *CalculatorTestSuite) TestCalculator_Sub() {
	suite.calcService.Num1 = 5
	suite.calcService.Num2 = 4
	actual := suite.calcService.Sub()
	expected := 1
	assert.Equal(suite.T(), expected, actual, "it should return 1")
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
