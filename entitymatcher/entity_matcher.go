package entitymatcher

import (
	"errors"
	"reflect"
	"unsafe"

	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/spf13/cast"
)

// EntityMatcher implements Matcher for matching
// filters against in-memory go objects
type EntityMatcher struct {
	inputs     interface{}
	conditions ConditionMap
}

// NewEntityMatcher constructs a new EntityMatcher
func NewEntityMatcher(inputs interface{}) EntityMatcher {
	return EntityMatcher{
		inputs:     inputs,
		conditions: NewConditionMap(),
	}
}

func (m EntityMatcher) matchesFilter(entity reflect.Value, filter filter.FilterEntry) (bool, error) {
	entityField := entity.FieldByName(filter.AttributeName)
	if !entityField.IsValid() {
		return false, errors.New("unable to find attribute value on input entity")
	}

	if filter.IsWildcard {
		return true, nil
	}

	entityField = reflect.NewAt(
		entityField.Type(),
		unsafe.Pointer(entityField.UnsafeAddr()),
	).Elem()

	entityValStr, err := cast.ToStringE(entityField.Interface())
	if err != nil {
		return false, err
	}

	return m.conditions.MatchesCondition(
		entityValStr,
		filter.Value,
		filter.Condition,
	)
}

func (m EntityMatcher) matchesFilters(entity reflect.Value, filters ...filter.FilterEntry) (bool, error) {
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

func (m EntityMatcher) matchSingleEntity(
	entity interface{},
	filters ...filter.FilterEntry,
) ([]interface{}, error) {

	reflectedEntity := reflect.ValueOf(entity)
	matches, err := m.matchesFilters(reflectedEntity, filters...)
	if err != nil {
		return nil, err
	}

	if !matches {
		return []interface{}{}, nil
	}

	return []interface{}{entity}, nil
}

// GetMatches returns all input entities matching the set of provided filters.
// Any errors during filter processing will cause further processing to halt.
func (m EntityMatcher) GetMatches(filters ...filter.FilterEntry) ([]interface{}, error) {

	inputType := reflect.TypeOf(m.inputs).Kind()
	if inputType != reflect.Slice {
		return m.matchSingleEntity(m.inputs, filters...)
	}

	inputSlice := reflect.ValueOf(m.inputs)
	results := make([]interface{}, 0)

	for i := 0; i < inputSlice.Len(); i++ {
		entity := inputSlice.Index(i)
		matched, err := m.matchesFilters(entity, filters...)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, entity.Interface())
		}
	}

	return results, nil
}
