package filter

import (
	"errors"
	"strings"
)

const FILTER_WILDCARD_CHAR string = "*"

func getFilterCondition(filterStr string, conditions []FilterCondition) (FilterCondition, error) {
	for _, condition := range conditions {
		if strings.Contains(filterStr, string(condition)) {
			return condition, nil
		}
	}

	return FILTER_EQUALS, errors.New("unable to identify filter condition")
}

func buildFilterEntry(filterStr string) (FilterEntry, error) {
	condition, err := getFilterCondition(filterStr, FilterConditions)
	if err != nil {
		return FilterEntry{}, err
	}

	filterParts := strings.Split(filterStr, string(condition))
	if len(filterParts) != 2 {
		return FilterEntry{},
			errors.New("invalid filter format: " + filterStr)
	}

	filterValue := filterParts[1]
	isWildcard := strings.Contains(filterValue, FILTER_WILDCARD_CHAR)

	filter := FilterEntry{
		AttributeName: filterParts[0],
		Value:         filterValue,
		Condition:     condition,
		IsWildcard:    isWildcard,
	}

	return filter, nil
}

// ParseFilterStr constructs an array of entity filters from a filter string
func ParseFilterStr(filterStr string) (EntityFilter, error) {
	filterParts := strings.Split(filterStr, ",")

	filter := NewEntityFilter()
	for _, filterStr := range filterParts {
		if strings.TrimSpace(filterStr) == "" {
			continue
		}

		newEntry, err := buildFilterEntry(filterStr)
		if err != nil {
			return filter, err
		}

		filter.Add(newEntry)
	}

	return filter, nil
}
