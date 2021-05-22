package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Assets holds the schema definition for the Assets entity.
type Assets struct {
	ent.Schema
}

// Fields of the Assets.
func (Assets) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.String("ext"),
		field.String("type"),
	}
}

// Edges of the Assets.
func (Assets) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("view", View.Type),
	}
}
