package filter

import (
	"regexp"
	"strings"
)

type EntityFilter []FilterEntry

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
