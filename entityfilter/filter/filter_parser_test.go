package filter_test

import (
	"testing"

	"github.com/kgoins/entityfilter/entityfilter/filter"
)

func TestFilterParser_SingleFilterSingleCondition(t *testing.T) {
	filterStr := "username:=myuser"

	filters, err := filter.ParseFilterStr(filterStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(filters) != 1 {
		t.Fatal("Unable to construct filter")
	}

	myFilter := filters[0]

	if myFilter.Condition != filter.FILTER_EQUALS {
		t.Fatal("Unable to construct filter condition")
	}

	if myFilter.AttributeName != "username" || myFilter.Value != "myuser" {
		t.Fatal("Unable to construct filter values")
	}
}

func TestFilterParser_InvalidFilterCondition(t *testing.T) {
	filterStr := "username$=myuser"

	_, err := filter.ParseFilterStr(filterStr)
	if err == nil {
		t.Fatal("Failed to error on invalid filter condition")
	}
}
