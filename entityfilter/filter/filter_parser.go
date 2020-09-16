package filter

import (
	"errors"
	"sort"
	"strings"
)

func getFilterCondition(filterStr string, conditions []FilterCondition) (FilterCondition, error) {
	for _, condition := range conditions {
		if strings.Contains(filterStr, string(condition)) {
			return condition, nil
		}
	}

	return FILTER_EQUALS, errors.New("Unable to identify filter condition")
}

func buildEntityFilter(filterStr string) (EntityFilter, error) {
	condition, err := getFilterCondition(filterStr, FilterConditions)
	if err != nil {
		return EntityFilter{}, err
	}

	filterParts := strings.Split(filterStr, string(condition))
	if len(filterParts) != 2 {
		return EntityFilter{},
			errors.New("Invalid filter format: " + filterStr)
	}

	filterValue := filterParts[1]
	isWildcard := filterValue == "*"

	filter := EntityFilter{
		AttributeName: filterParts[0],
		Value:         filterValue,
		Condition:     condition,
		IsWildcard:    isWildcard,
	}

	return filter, nil
}

func BuildEntityFilter(filterStrings []string) ([]EntityFilter, error) {
	filter := []EntityFilter{}
	sort.Sort(filterByLength(FilterConditions))

	for _, filterStr := range filterStrings {
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