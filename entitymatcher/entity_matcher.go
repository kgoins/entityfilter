package entitymatcher

import (
	"errors"
	"reflect"

	filter "github.com/kgoins/entityfilter/entityfilter"
)

// EntityMatcher implements Matcher for matching
// filters against in-memory go objects
type EntityMatcher struct {
	extractor  FieldExtractor
	conditions ConditionMap
}

// NewEntityMatcher constructs a new EntityMatcher
func NewEntityMatcher(extractor FieldExtractor) EntityMatcher {
	return EntityMatcher{
		extractor:  extractor,
		conditions: NewConditionMap(),
	}
}

// NewStructEntityMatcher constructs a new EntityMatcher with
// a default struct field extractor (works for most structs)
func NewStructEntityMatcher() EntityMatcher {
	return EntityMatcher{
		extractor:  NewStructFieldExtractor(),
		conditions: NewConditionMap(),
	}
}

func (m EntityMatcher) matchesFilter(entity interface{}, filter filter.FilterEntry) (bool, error) {
	val, err := m.extractor.GetFieldValue(entity, filter.AttributeName)
	if err != nil {
		return false, err
	}

	if filter.IsWildcard {
		return true, nil
	}

	return m.conditions.MatchesCondition(
		val,
		filter.Value,
		filter.Condition,
	)
}

func (m EntityMatcher) matchesFilters(entity interface{}, filters ...filter.FilterEntry) (bool, error) {
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

func (m EntityMatcher) Matches(entity interface{}, filter filter.EntityFilter) (bool, error) {
	return m.matchesFilters(entity, filter.GetEntries()...)
}

// GetMatches returns all input entities matching the set of provided filters.
// Any errors during filter processing will cause further processing to halt.
func (m EntityMatcher) GetMatches(
	inputs interface{},
	filter filter.EntityFilter,
) ([]interface{}, error) {

	inputType := reflect.TypeOf(inputs).Kind()
	if inputType != reflect.Slice {
		return nil, errors.New("input type is not a slice")
	}

	inputSlice := reflect.ValueOf(inputs)
	inputStream := make(chan interface{})

	go func() {
		defer close(inputStream)
		for i := 0; i < inputSlice.Len(); i++ {
			entity := inputSlice.Index(i).Interface()
			inputStream <- entity
		}
	}()

	resultStream := m.GetMatchesStream(inputStream, filter)

	results := []interface{}{}
	for r := range resultStream {
		results = append(results, r)
	}

	return results, nil
}

func (m EntityMatcher) GetMatchesStream(
	inputs <-chan interface{},
	filter filter.EntityFilter,
) <-chan interface{} {

	results := make(chan interface{})

	go func() {
		defer close(results)
		for entity := range inputs {
			matched, err := m.matchesFilters(entity, filter.GetEntries()...)
			if err != nil {
				continue
			}

			if matched {
				results <- entity
			}
		}
	}()

	return results
}
