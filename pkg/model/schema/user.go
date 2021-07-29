package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/maxiloEmmmm/go-web/contact"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(255).Unique(),
		field.String("password").MaxLen(255),
		field.Int8("enable").GoType(&contact.BoolField{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
