package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SubItem holds the schema definition for the SubItem entity.
type SubItem struct {
	ent.Schema
}

// Fields of the SubItem.
func (SubItem) Fields() []ent.Field {
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
	}
}

// Edges of the SubItem.
func (SubItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", BagItem.Type).
			Ref("sub_items").
			Unique(),
	}
}
