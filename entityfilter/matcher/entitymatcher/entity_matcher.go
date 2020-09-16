package entitymatcher

import (
	"errors"
	"reflect"

	"github.com/kgoins/entityfilter/entityfilter/filter"
	"github.com/spf13/cast"
)

// EntityMatcher implements Matcher for matching
// filters against in-memory go objects
type EntityMatcher struct {
	inputs     []interface{}
	conditions ConditionMap
}

// NewEntityMatcher constructs a new EntityMatcher
func NewEntityMatcher(inputs []interface{}) EntityMatcher {
	return EntityMatcher{
		inputs:     inputs,
		conditions: NewConditionMap(),
	}
}

func (m EntityMatcher) matchesFilter(entity interface{}, filter filter.EntityFilter) (bool, error) {
	reflectedEntity := reflect.Indirect(reflect.ValueOf(entity))

	entityVal := reflectedEntity.FieldByName(filter.AttributeName)
	if !entityVal.IsValid() {
		return false, errors.New("unable to find attribute value on input entity")
	}

	if filter.IsWildcard {
		return true, nil
	}

	entityValStr, err := cast.ToStringE(entityVal)
	if err != nil {
		return false, err
	}

	return m.conditions.MatchesCondition(
		entityValStr,
		filter.Value,
		filter.Condition,
	)
}

func (m EntityMatcher) matchesFilters(entity interface{}, filters []filter.EntityFilter) (bool, error) {
	for _, filter := range filters {
		matched, err := m.matchesFilter(entity, filter)
		if err != nil {
			return false, err
		}

		if !matched {
			return false, nil
		}
	}

	return true, nil
}

// GetMatches returns all input entities matching the set of provided filters.
// Any errors during filter processing will cause further processing to halt.
func (m EntityMatcher) GetMatches(filters []filter.EntityFilter) ([]interface{}, error) {
	results := make([]interface{}, 0)

	for _, entity := range m.inputs {
		matched, err := m.matchesFilters(entity, filters)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, entity)
		}
	}

	return results, nil
}
