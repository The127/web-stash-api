package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BagItem holds the schema definition for the BagItem entity.
type BagItem struct {
	ent.Schema
}

// Fields of the BagItem.
func (BagItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.String("description").
			Optional(),
		field.String("icon").
			NotEmpty(),
		field.String("link").
			NotEmpty(),
		field.String("image").
			NotEmpty(),
	}
}

// Edges of the BagItem.
func (BagItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bag", Bag.Type).
			Ref("items").
			Unique(),
		edge.To("sub_items", SubItem.Type),
	}
}
