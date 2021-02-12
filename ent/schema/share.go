package schema

import "entgo.io/ent"

// Share holds the schema definition for the Share entity.
type Share struct {
	ent.Schema
}

// Fields of the Share.
func (Share) Fields() []ent.Field {
	return nil
}

// Edges of the Share.
func (Share) Edges() []ent.Edge {
	return nil
}
