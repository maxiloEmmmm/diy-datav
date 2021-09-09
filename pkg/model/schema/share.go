package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Share holds the schema definition for the Share entity.
type Share struct {
	ent.Schema
}

// Fields of the Share.
func (Share) Fields() []ent.Field {
	return []ent.Field{
		field.Time("end_at"),
	}
}

// Edges of the Share.
func (Share) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("view", View.Type).Ref("share").Unique(),
		edge.From("creator", User.Type).Ref("share").Unique(),
	}
}
