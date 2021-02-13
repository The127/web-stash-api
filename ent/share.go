// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"web-stash-api/ent/share"

	"github.com/google/uuid"
)

// Share is the model entity for the Share schema.
type Share struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// ItemID holds the value of the "item_id" field.
	ItemID uuid.UUID `json:"item_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Share) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case share.FieldID, share.FieldUserID, share.FieldItemID:
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Share", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Share fields.
func (s *Share) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case share.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case share.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				s.UserID = *value
			}
		case share.FieldItemID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field item_id", values[i])
			} else if value != nil {
				s.ItemID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Share.
// Note that you need to call Share.Unwrap() before calling this method if this Share
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Share) Update() *ShareUpdateOne {
	return (&ShareClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Share entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Share) Unwrap() *Share {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Share is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Share) String() string {
	var builder strings.Builder
	builder.WriteString("Share(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", item_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ItemID))
	builder.WriteByte(')')
	return builder.String()
}

// Shares is a parsable slice of Share.
type Shares []*Share

func (s Shares) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
