package filter

type EntityFilter struct {
	AttributeName string
	Value         string
	Condition     FilterCondition
	IsWildcard    bool
}
