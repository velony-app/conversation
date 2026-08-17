package value_object

import (
	"errors"

	"go.einride.tech/aip/resourcename"
)

var ErrResourceNameInvalid = errors.New("resource name is invalid")

type ResourceName struct {
	value string
}

func NewResourceName(value string) (ResourceName, error) {
	if value == "" {
		return ResourceName{}, nil
	}

	if err := resourcename.Validate(value); err != nil {
		return ResourceName{}, ErrResourceNameInvalid
	}

	return ResourceName{
		value: value,
	}, nil
}

func (r ResourceName) Value() string {
	return r.value
}

func (r ResourceName) String() string {
	return r.value
}

func (r ResourceName) IsEmpty() bool {
	return r.value == ""
}
