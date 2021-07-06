package filter

import (
	"regexp"
	"strings"
)

type EntityFilter struct {
	AttributeName string
	Value         string
	Condition     FilterCondition
	IsWildcard    bool
}

func (ef EntityFilter) BuildWildcardRegexPattern() string {
	return strings.ReplaceAll(ef.Value, FILTER_WILDCARD_CHAR, ".*")
}

func (ef EntityFilter) BuildWildcardRegex() (*regexp.Regexp, error) {
	regexStr := ef.BuildWildcardRegexPattern()
	filterRegex, err := regexp.Compile(regexStr)
	if err != nil {
		return nil, err
	}

	return filterRegex, nil
}
