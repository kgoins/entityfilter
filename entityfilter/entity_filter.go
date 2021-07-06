package filter

import (
	"regexp"
	"strings"
)

type EntityFilter struct {
	filter []FilterEntry
}

func NewEntityFilter(entries ...FilterEntry) EntityFilter {
	filter := []FilterEntry{}
	filter = append(filter, entries...)
	return EntityFilter{filter: filter}
}

func (f *EntityFilter) Add(entry FilterEntry) {
	f.filter = append(f.filter, entry)
}

func (f EntityFilter) GetEntries() []FilterEntry {
	return f.filter
}

func (f EntityFilter) Len() int {
	return len(f.filter)
}

type FilterEntry struct {
	AttributeName string
	Value         string
	Condition     FilterCondition
	IsWildcard    bool
}

func (ef FilterEntry) BuildWildcardRegexPattern() string {
	return strings.ReplaceAll(ef.Value, FILTER_WILDCARD_CHAR, ".*")
}

func (ef FilterEntry) BuildWildcardRegex() (*regexp.Regexp, error) {
	regexStr := ef.BuildWildcardRegexPattern()
	filterRegex, err := regexp.Compile(regexStr)
	if err != nil {
		return nil, err
	}

	return filterRegex, nil
}
