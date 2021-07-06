package entitymatcher_test

import (
	"testing"

	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/entityfilter/entitymatcher"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	MyInt       int
	MyStr       string
	secretField string
}

func getTestStructs() []TestStruct {
	return []TestStruct{
		{MyInt: 0, MyStr: "hello"},
		{MyInt: 1, MyStr: "hello world"},
		{MyInt: 2, MyStr: "Hello World!"},
	}
}

func TestEntityMatcher_MatchSingleEntity(t *testing.T) {
	r := require.New(t)

	filterEntry := filter.FilterEntry{
		AttributeName: "MyInt",
		Value:         "1",
		Condition:     filter.FILTER_EQUALS,
		IsWildcard:    false,
	}

	filter := filter.NewEntityFilter(filterEntry)
	testStruct := TestStruct{MyInt: 1, MyStr: "hello"}

	matcher := entitymatcher.NewStructEntityMatcher()
	matches, err := matcher.Matches(testStruct, filter)
	r.NoError(err)
	r.True(matches)
}

func TestEntityMatcher_UnexportedField(t *testing.T) {
	r := require.New(t)

	filterEntry := filter.FilterEntry{
		AttributeName: "secretField",
		Value:         "x",
		Condition:     filter.FILTER_EQUALS,
		IsWildcard:    false,
	}

	filter := filter.NewEntityFilter(filterEntry)
	testStruct := TestStruct{MyInt: 1, MyStr: "hello", secretField: "x"}

	matcher := entitymatcher.NewStructEntityMatcher()
	_, err := matcher.Matches(testStruct, filter)
	r.Error(err)
}

func TestEntityMatcher_MatchSingleFilterMultEntities(t *testing.T) {
	r := require.New(t)

	filterEntry := filter.FilterEntry{
		AttributeName: "MyInt",
		Value:         "1",
		Condition:     filter.FILTER_EQUALS,
		IsWildcard:    false,
	}

	filter := filter.NewEntityFilter(filterEntry)
	testStructs := getTestStructs()

	matcher := entitymatcher.NewStructEntityMatcher()
	results, err := matcher.GetMatches(testStructs, filter)
	r.NoError(err)
	r.Equal(1, len(results))
}

func TestEntityMatcher_CompositeFilterMultResults(t *testing.T) {
	r := require.New(t)

	testFilterStr := "MyInt:!=0,MyStr:~hello"
	testStructs := getTestStructs()

	testFilter, err := filter.ParseFilterStr(testFilterStr)
	r.NoError(err)

	matcher := entitymatcher.NewStructEntityMatcher()
	results, err := matcher.GetMatches(testStructs, testFilter)
	r.NoError(err)
	r.Equal(1, len(results))
}
