package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MyTestSuiteStruct struct {
	suite.Suite
}

func (suite *MyTestSuiteStruct) BeforeTest(suiteName string, testName string) {
	suite.T().Log("BeforeTest!!")
}

func (suite *MyTestSuiteStruct) AfterTest(suiteName string, testName string) {
	suite.T().Log("AfterTest!!")
}

func (suite *MyTestSuiteStruct) TestHello() {
	assert.Equal(suite.T(), 1, 1)
}

func (suite *MyTestSuiteStruct) TestBool() {
	assert.True(suite.T(), false)
}

func TestFirstTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuiteStruct))
}
