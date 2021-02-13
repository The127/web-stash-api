package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Share holds the schema definition for the Share entity.
type Share struct {
	ent.Schema
}

// Fields of the Share.
func (Share) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).
			Immutable(),
		field.UUID("item_id", uuid.UUID{}).
			Immutable(),
	}
}

func (Share) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "item_id").
			Unique(),
	}
}

// Edges of the Share.
func (Share) Edges() []ent.Edge {
	return []ent.Edge{}
}
