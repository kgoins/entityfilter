package entitymatcher

import "github.com/kgoins/entityfilter/entityfilter/filter"

// EntityMatcher implements Matcher for matching
// filters against in-memory go objects
type EntityMatcher struct {
	inputs []interface{}
}

// NewEntityMatcher constructs a new EntityMatcher
func NewEntityMatcher(inputs []interface{}) EntityMatcher {
	return EntityMatcher{
		inputs: inputs,
	}
}

func (m EntityMatcher) GetMatches(filters []filter.EntityFilter) ([]interface{}, error) {
	return nil, nil
}
