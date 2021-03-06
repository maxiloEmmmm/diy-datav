// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEnable holds the string denoting the enable field in the database.
	FieldEnable = "enable"
	// EdgeShare holds the string denoting the share edge name in mutations.
	EdgeShare = "share"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ShareTable is the table the holds the share relation/edge.
	ShareTable = "shares"
	// ShareInverseTable is the table name for the Share entity.
	// It exists in this package in order to avoid circular dependency with the "share" package.
	ShareInverseTable = "shares"
	// ShareColumn is the table column denoting the share relation/edge.
	ShareColumn = "user_share"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldEnable,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
