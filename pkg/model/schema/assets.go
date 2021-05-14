package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Assets holds the schema definition for the Assets entity.
type Assets struct {
	ent.Schema
}

// Fields of the Assets.
func (Assets) Fields() []ent.Field {
	return nil
}

// Edges of the Assets.
func (Assets) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("view", View.Type),
	}
}
