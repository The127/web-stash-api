// Code generated by entc, DO NOT EDIT.

package bagitem

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the bagitem type in the database.
	Label = "bag_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"

	// EdgeBag holds the string denoting the bag edge name in mutations.
	EdgeBag = "bag"
	// EdgeSubItems holds the string denoting the sub_items edge name in mutations.
	EdgeSubItems = "sub_items"

	// Table holds the table name of the bagitem in the database.
	Table = "bag_items"
	// BagTable is the table the holds the bag relation/edge.
	BagTable = "bag_items"
	// BagInverseTable is the table name for the Bag entity.
	// It exists in this package in order to avoid circular dependency with the "bag" package.
	BagInverseTable = "bags"
	// BagColumn is the table column denoting the bag relation/edge.
	BagColumn = "bag_items"
	// SubItemsTable is the table the holds the sub_items relation/edge.
	SubItemsTable = "sub_items"
	// SubItemsInverseTable is the table name for the SubItem entity.
	// It exists in this package in order to avoid circular dependency with the "subitem" package.
	SubItemsInverseTable = "sub_items"
	// SubItemsColumn is the table column denoting the sub_items relation/edge.
	SubItemsColumn = "bag_item_sub_items"
)

// Columns holds all SQL columns for bagitem fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldIcon,
	FieldLink,
	FieldImage,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the BagItem type.
var ForeignKeys = []string{
	"bag_items",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// IconValidator is a validator for the "icon" field. It is called by the builders before save.
	IconValidator func(string) error
	// LinkValidator is a validator for the "link" field. It is called by the builders before save.
	LinkValidator func(string) error
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
