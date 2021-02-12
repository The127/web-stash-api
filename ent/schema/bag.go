package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Bag holds the schema definition for the Bag entity.
type Bag struct {
	ent.Schema
}

// Fields of the Bag.
func (Bag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.String("icon").
			MaxLen(255).
			NotEmpty(),
		field.UUID("user_id", uuid.UUID{}),
	}
}

func (Bag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "name").
			Unique(),
	}
}

// Edges of the Bag.
func (Bag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", BagItem.Type),
	}
}
