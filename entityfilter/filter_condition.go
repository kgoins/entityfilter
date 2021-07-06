package filter

type FilterCondition string

const (
	FILTER_EQUALS       FilterCondition = ":="
	FILTER_NOT_EQUALS   FilterCondition = ":!="
	FILTER_CONTAINS     FilterCondition = ":~"
	FILTER_NOT_CONTAINS FilterCondition = ":!~"
)

var FilterConditions = []FilterCondition{
	FILTER_EQUALS,
	FILTER_NOT_EQUALS,
	FILTER_CONTAINS,
	FILTER_NOT_CONTAINS,
}

type filterByLength []FilterCondition

func (a filterByLength) Len() int      { return len(a) }
func (a filterByLength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a filterByLength) Less(i, j int) bool {
	li, lj := len(a[i]), len(a[j])
	if li == lj {
		return a[i] > a[j]
	}
	return li > lj
}
