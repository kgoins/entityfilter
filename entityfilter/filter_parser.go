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

func buildEntityFilter(filterStr string) (EntityFilter, error) {
	condition, err := getFilterCondition(filterStr, FilterConditions)
	if err != nil {
		return EntityFilter{}, err
	}

	filterParts := strings.Split(filterStr, string(condition))
	if len(filterParts) != 2 {
		return EntityFilter{},
			errors.New("invalid filter format: " + filterStr)
	}

	filterValue := filterParts[1]
	isWildcard := strings.Contains(filterValue, FILTER_WILDCARD_CHAR)

	filter := EntityFilter{
		AttributeName: filterParts[0],
		Value:         filterValue,
		Condition:     condition,
		IsWildcard:    isWildcard,
	}

	return filter, nil
}

// ParseFilterStr constructs an array of entity filters from a filter string
func ParseFilterStr(filterStr string) ([]EntityFilter, error) {
	filterParts := strings.Split(filterStr, ",")

	filter := []EntityFilter{}
	for _, filterStr := range filterParts {
		if strings.TrimSpace(filterStr) == "" {
			continue
		}

		newFilter, err := buildEntityFilter(filterStr)
		if err != nil {
			return nil, err
		}

		filter = append(filter, newFilter)
	}

	return filter, nil
}
