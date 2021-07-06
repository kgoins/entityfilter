package entitymatcher

import (
	"errors"
	"strings"

	filter "github.com/kgoins/entityfilter/entityfilter"
)

type FilterFunc func(string, string) (bool, error)

type ConditionMap struct {
	conditions map[filter.FilterCondition]FilterFunc
}

func (m ConditionMap) MatchesCondition(
	entityValue string,
	filterValue string,
	condition filter.FilterCondition,
) (bool, error) {
	filterFunc, found := m.conditions[condition]
	if !found {
		return false, errors.New("condition not supported")
	}

	return filterFunc(entityValue, filterValue)
}

func NewConditionMap() ConditionMap {
	condMap := make(map[filter.FilterCondition]FilterFunc)

	condMap[filter.FILTER_EQUALS] = filterEquals
	condMap[filter.FILTER_CONTAINS] = filterContains
	condMap[filter.FILTER_NOT_EQUALS] = filterNotEquals
	condMap[filter.FILTER_NOT_CONTAINS] = filterNotContains

	return ConditionMap{conditions: condMap}
}

func filterEquals(entityValue string, filterValue string) (bool, error) {
	return entityValue == filterValue, nil
}

func filterContains(entityValue string, filterValue string) (bool, error) {
	return strings.Contains(entityValue, filterValue), nil
}

func filterNotEquals(entityValue string, filterValue string) (bool, error) {
	return entityValue != filterValue, nil
}

func filterNotContains(entityValue string, filterValue string) (bool, error) {
	return !strings.Contains(entityValue, filterValue), nil
}
