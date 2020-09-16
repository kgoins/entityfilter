package entitymatcher_test

import (
	"testing"

	"github.com/kgoins/entityfilter/entityfilter/filter"
	"github.com/kgoins/entityfilter/entityfilter/matcher/entitymatcher"
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
