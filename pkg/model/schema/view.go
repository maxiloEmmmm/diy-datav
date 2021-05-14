package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// View holds the schema definition for the View entity.
type View struct {
	ent.Schema
}

// Fields of the View.
func (View) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc").MaxLen(255),
		field.Text("config"),
	}
}

// Edges of the View.
func (View) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bg", Assets.Type).Ref("view").Unique(),
		edge.To("blocks", ViewBlock.Type),
	}
}
