package entitymatcher_test

import (
	"testing"

	"gitlab.rule10.net/rule10/entityfilter/entityfilter/filter"
	"gitlab.rule10.net/rule10/entityfilter/entityfilter/matcher/entitymatcher"
)

type TestStruct struct {
	myint int
	mystr string
}

func TestEntityMatcher_MatchSingleFilterMultTypes(t *testing.T) {
	testFilter := filter.EntityFilter{
		AttributeName: "myint",
		Value:         "1",
		Condition:     filter.FILTER_EQUALS,
		IsWildcard:    false,
	}

	testStructs := []TestStruct{
		TestStruct{myint: 0, mystr: "hello"},
		TestStruct{myint: 1, mystr: "hello world"},
		TestStruct{myint: 2, mystr: "Hello World!"},
	}

	matcher := entitymatcher.NewEntityMatcher(testStructs)
	results, err := matcher.GetMatches(testFilter)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(results) != 1 {
		t.Fatal("Wrong number of matches returned")
	}
}

func TestEntityMatcher_CompositeFilterMultResults(t *testing.T) {
	testFilterStr := "myint:!=0,mystr:~hello"

	testStructs := []TestStruct{
		TestStruct{myint: 0, mystr: "hello"},
		TestStruct{myint: 1, mystr: "hello world"},
		TestStruct{myint: 2, mystr: "hello world!"},
	}

	testFilter, err := filter.ParseFilterStr(testFilterStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	matcher := entitymatcher.NewEntityMatcher(testStructs)
	results, err := matcher.GetMatches(testFilter...)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(results) != 2 {
		t.Fatal("Wrong number of matches returned")
	}
}
