package value_object

import "github.com/google/uuid"

type UserID struct {
	value uuid.UUID
}

func NewUserID(value uuid.UUID) UserID {
	return UserID{
		value: value,
	}
}

func (id UserID) Value() uuid.UUID {
	return id.value
}

func (id UserID) String() string {
	return id.value.String()
}
