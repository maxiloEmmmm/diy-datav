package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TypeConfig holds the schema definition for the TypeConfig entity.
type TypeConfig struct {
	ent.Schema
}

// Fields of the TypeConfig.
func (TypeConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("type").MaxLen(255),
		field.Text("config"),
	}
}

// Edges of the TypeConfig.
func (TypeConfig) Edges() []ent.Edge {
	return nil
}
