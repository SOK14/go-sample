package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SampleTestStruct struct {
	suite.Suite

	ExecutedTestMethodNames []string

	CallSetupSuiteCount    int
	CallSetupTestCount     int
	CallBeforeTestCounts   map[string]int
	runTestMethodNames     []string
	CallAfterTestCounts    map[string]int
	CallTearDownTestCount  int
	CallTearDownSuiteCount int
}

func (suite *SampleTestStruct) SetupSuite() {
	suite.T().Logf("===== SetupSuite =====")
	suite.CallSetupSuiteCount += 1
}

func (suite *SampleTestStruct) SetupTest() {
	suite.T().Logf("===== SetupTest =====")
	suite.CallSetupTestCount += 1
}

func (suite *SampleTestStruct) BeforeTest(suiteName string, testName string) {
	suite.T().Logf("===== BeforeTest suite[%s] test[%s] =====", suiteName, testName)

	if suite.CallBeforeTestCounts == nil {
		suite.CallBeforeTestCounts = make(map[string]int)
	}

	suite.CallBeforeTestCounts[suiteName+"/"+testName] += 1
}

func (suite *SampleTestStruct) AfterTest(suiteName string, testName string) {
	suite.T().Logf("===== AfterTest suite[%s] test[%s] =====", suiteName, testName)

	if suite.CallAfterTestCounts == nil {
		suite.CallAfterTestCounts = make(map[string]int)
	}

	suite.CallAfterTestCounts[suiteName+"/"+testName] += 1
}

func (suite *SampleTestStruct) TearDownTest() {
	suite.T().Logf("===== TearDownTest =====")
	suite.CallTearDownTestCount += 1
}

func (suite *SampleTestStruct) TearDownSuite() {
	suite.T().Logf("===== TearDownSuite =====")
	suite.CallTearDownSuiteCount += 1

	beforeTestCounts, _ := json.Marshal(suite.CallBeforeTestCounts)
	afterTestCounts, _ := json.Marshal(suite.CallAfterTestCounts)
	runTestMethodNames, _ := json.Marshal(suite.runTestMethodNames)

	fmt.Printf(`----------------------------------------------------------------------
  SetupSuiteCount = %d
  SetupTestCount = %d
  BeforeTestCounts = (
    %s
  )
  runTestMethods = %s
  AfterTestCounts = (
    %s
  )
  TearDownTestCount = %d
  TearDownSuiteCount = %d
----------------------------------------------------------------------
`,
		suite.CallSetupSuiteCount,
		suite.CallSetupTestCount,
		beforeTestCounts,
		runTestMethodNames,
		afterTestCounts,
		suite.CallTearDownTestCount,
		suite.CallTearDownSuiteCount,
	)
}

func (suite *SampleTestStruct) TestPlus() {
	suite.T().Logf("run %s", suite.T().Name())

	if suite.runTestMethodNames == nil {
		suite.runTestMethodNames = make([]string, 0)
	}

	suite.runTestMethodNames = append(suite.runTestMethodNames, suite.T().Name())

	assert.Equal(suite.T(), 1, 1)
}

func (suite *SampleTestStruct) TestBool() {
	suite.T().Logf("run %s", suite.T().Name())

	if suite.runTestMethodNames == nil {
		suite.runTestMethodNames = make([]string, 0)
	}

	suite.runTestMethodNames = append(suite.runTestMethodNames, suite.T().Name())

	assert.True(suite.T(), true)
}

func TestSuiteExample(t *testing.T) {
	suite.Run(t, new(SampleTestStruct))
}
