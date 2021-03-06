// Code generated by entc, DO NOT EDIT.

package menu

import (
	"time"
)

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the menu in the database.
	Table = "menus"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "menus"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "menu_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "menus"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "menu_children"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldURL,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "menus"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"menu_children",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
