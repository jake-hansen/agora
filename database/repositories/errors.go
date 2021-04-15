package repositories

import "fmt"

const (
	DATABASE_ACTION_RETRIEVE = "retrieving"
	DATABASE_ACTION_CREATE = "creating"
	DATABASE_ACTION_DELETE = "deleting"
)

type DatabaseError struct {
	DatabaseAction string
	EntityType string
	EntityValue string
	EntityField string
	NestedError error
}

func (d DatabaseError) Error() string {
	return fmt.Sprintf("error %s %s %s %s: %s", d.DatabaseAction, d.EntityType, d.EntityField, d.EntityValue, d.NestedError.Error())
}

func (d DatabaseError) Unwrap() error {
	return d.NestedError
}

func NewDatabaseError(action string, entityType string, entityValue string, entityField string, nestedError error) DatabaseError {
	return DatabaseError{
		DatabaseAction: action,
		EntityType:     entityType,
		EntityValue:    entityValue,
		EntityField:    entityField,
		NestedError:    nestedError,
	}
}

type NotFoundError struct {
	Value string
}

func (n NotFoundError) Error() string {
	return "record not found"
}

func NewNotFoundError(action string, entityType string, entityValue string, entityField string) DatabaseError {
	return DatabaseError{
		DatabaseAction: action,
		EntityType:     entityType,
		EntityValue:    entityValue,
		EntityField:    entityField,
		NestedError:    NotFoundError{
			Value: entityValue,
		},
	}
}



