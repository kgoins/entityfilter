package filter_test

import (
	"testing"

	filter "github.com/kgoins/entityfilter/entityfilter"
)

func TestFilterParser_SingleFilterSingleCondition(t *testing.T) {
	filterStr := "username:=myuser"

	filters, err := filter.ParseFilterStr(filterStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if filters.Len() != 1 {
		t.Fatal("Unable to construct filter")
	}

	myFilter := filters.GetEntries()[0]

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

func TestFilterParser_CompositeFilter(t *testing.T) {
	filterStr := "username:=myuser,description:~vpn,myfield:!=thisvalue"

	filters, err := filter.ParseFilterStr(filterStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if filters.Len() != 3 {
		t.Fatal("Unable to construct composite filter")
	}
}

func TestFilterParser_CompositeWithWildcard(t *testing.T) {
	filterStr := "username:=my*user,description:~vpn,myfield:!=thisvalue*"

	filters, err := filter.ParseFilterStr(filterStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	targetFilter := filters.GetEntries()[0]
	if !targetFilter.IsWildcard {
		t.Fatal("Unable to identify wildcard in filter str")
	}

	targetFilter = filters.GetEntries()[2]
	if !targetFilter.IsWildcard {
		t.Fatal("Unable to identify wildcard in filter str")
	}
}
