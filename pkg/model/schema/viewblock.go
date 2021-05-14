package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ViewBlock holds the schema definition for the ViewBlock entity.
type ViewBlock struct {
	ent.Schema
}

// Fields of the ViewBlock.
func (ViewBlock) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").MaxLen(255),
		field.Text("config"),
	}
}

// Edges of the ViewBlock.
func (ViewBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("view", View.Type).Ref("blocks").Unique(),
		edge.To("dataset", DataSet.Type),
	}
}
