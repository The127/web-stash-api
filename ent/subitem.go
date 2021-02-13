// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"web-stash-api/ent/bagitem"
	"web-stash-api/ent/subitem"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// SubItem is the model entity for the SubItem schema.
type SubItem struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubItemQuery when eager-loading is set.
	Edges              SubItemEdges `json:"edges"`
	bag_item_sub_items *uuid.UUID
}

// SubItemEdges holds the relations/edges for other nodes in the graph.
type SubItemEdges struct {
	// Parent holds the value of the parent edge.
	Parent *BagItem `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubItemEdges) ParentOrErr() (*BagItem, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bagitem.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case subitem.FieldName, subitem.FieldDescription, subitem.FieldIcon, subitem.FieldLink:
			values[i] = &sql.NullString{}
		case subitem.FieldID:
			values[i] = &uuid.UUID{}
		case subitem.ForeignKeys[0]: // bag_item_sub_items
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type SubItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubItem fields.
func (si *SubItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subitem.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				si.ID = *value
			}
		case subitem.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				si.Name = value.String
			}
		case subitem.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				si.Description = value.String
			}
		case subitem.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				si.Icon = value.String
			}
		case subitem.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				si.Link = value.String
			}
		case subitem.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field bag_item_sub_items", values[i])
			} else if value != nil {
				si.bag_item_sub_items = value
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the SubItem entity.
func (si *SubItem) QueryParent() *BagItemQuery {
	return (&SubItemClient{config: si.config}).QueryParent(si)
}

// Update returns a builder for updating this SubItem.
// Note that you need to call SubItem.Unwrap() before calling this method if this SubItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (si *SubItem) Update() *SubItemUpdateOne {
	return (&SubItemClient{config: si.config}).UpdateOne(si)
}

// Unwrap unwraps the SubItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (si *SubItem) Unwrap() *SubItem {
	tx, ok := si.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubItem is not a transactional entity")
	}
	si.config.driver = tx.drv
	return si
}

// String implements the fmt.Stringer.
func (si *SubItem) String() string {
	var builder strings.Builder
	builder.WriteString("SubItem(")
	builder.WriteString(fmt.Sprintf("id=%v", si.ID))
	builder.WriteString(", name=")
	builder.WriteString(si.Name)
	builder.WriteString(", description=")
	builder.WriteString(si.Description)
	builder.WriteString(", icon=")
	builder.WriteString(si.Icon)
	builder.WriteString(", link=")
	builder.WriteString(si.Link)
	builder.WriteByte(')')
	return builder.String()
}

// SubItems is a parsable slice of SubItem.
type SubItems []*SubItem

func (si SubItems) config(cfg config) {
	for _i := range si {
		si[_i].config = cfg
	}
}
