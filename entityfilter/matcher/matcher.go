package matcher

import "github.com/kgoins/entityfilter/entityfilter/filter"

// Matcher represents the ability to determine if a given data
// source contains any entities matching the input filter conditions.
// Data sources are configured by individual implementations during
// construction. Input filters are run against that data source.
type Matcher interface {
	GetMatches(...filter.EntityFilter) ([]interface{}, error)
}
