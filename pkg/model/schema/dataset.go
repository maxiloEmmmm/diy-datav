package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DataSet holds the schema definition for the DataSet entity.
type DataSet struct {
	ent.Schema
}

// Fields of the DataSet.
func (DataSet) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").MaxLen(255),
		field.Text("title"),
		field.Text("config"),
	}
}

// Edges of the DataSet.
func (DataSet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("block", ViewBlock.Type).Ref("dataset").Unique(),
	}
}
