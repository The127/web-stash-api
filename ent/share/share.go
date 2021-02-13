// Code generated by entc, DO NOT EDIT.

package share

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the share type in the database.
	Label = "share"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldItemID holds the string denoting the item_id field in the database.
	FieldItemID = "item_id"

	// Table holds the table name of the share in the database.
	Table = "shares"
)

// Columns holds all SQL columns for share fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldItemID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)