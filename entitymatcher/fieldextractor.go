package entitymatcher

import (
	"errors"

	"github.com/fatih/structs"
	"github.com/spf13/cast"
)

type FieldExtractor interface {
	GetFieldValue(entity interface{}, fieldName string) (string, error)
}

type StructFieldExtractor struct{}

func NewStructFieldExtractor() StructFieldExtractor {
	return StructFieldExtractor{}
}

func (e StructFieldExtractor) GetFieldValue(entity interface{}, fieldName string) (string, error) {
	s := structs.New(entity)
	f, ok := s.FieldOk(fieldName)
	if !ok {
		return "", errors.New(
			"unable to find field: " + fieldName,
		)
	}

	if !f.IsExported() {
		return "", errors.New(
			"unable to access unexported field: " + fieldName,
		)
	}

	val := f.Value()
	return cast.ToStringE(val)
}
