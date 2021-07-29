// Code generated by entc, DO NOT EDIT.

package model

import (
	"time"

	"github.com/maxiloEmmmm/diy-datav/pkg/model/dataset"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/menu"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/schema"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/user"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	datasetFields := schema.DataSet{}.Fields()
	_ = datasetFields
	// datasetDescType is the schema descriptor for type field.
	datasetDescType := datasetFields[0].Descriptor()
	// dataset.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	dataset.TypeValidator = datasetDescType.Validators[0].(func(string) error)
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescCreatedAt is the schema descriptor for created_at field.
	menuDescCreatedAt := menuFields[2].Descriptor()
	// menu.DefaultCreatedAt holds the default value on creation for the created_at field.
	menu.DefaultCreatedAt = menuDescCreatedAt.Default.(func() time.Time)
	// menuDescUpdatedAt is the schema descriptor for updated_at field.
	menuDescUpdatedAt := menuFields[3].Descriptor()
	// menu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menu.DefaultUpdatedAt = menuDescUpdatedAt.Default.(func() time.Time)
	// menu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menu.UpdateDefaultUpdatedAt = menuDescUpdatedAt.UpdateDefault.(func() time.Time)
	typeconfigFields := schema.TypeConfig{}.Fields()
	_ = typeconfigFields
	// typeconfigDescType is the schema descriptor for type field.
	typeconfigDescType := typeconfigFields[1].Descriptor()
	// typeconfig.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	typeconfig.TypeValidator = typeconfigDescType.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	viewFields := schema.View{}.Fields()
	_ = viewFields
	// viewDescDesc is the schema descriptor for desc field.
	viewDescDesc := viewFields[0].Descriptor()
	// view.DescValidator is a validator for the "desc" field. It is called by the builders before save.
	view.DescValidator = viewDescDesc.Validators[0].(func(string) error)
	viewblockFields := schema.ViewBlock{}.Fields()
	_ = viewblockFields
	// viewblockDescType is the schema descriptor for type field.
	viewblockDescType := viewblockFields[0].Descriptor()
	// viewblock.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	viewblock.TypeValidator = viewblockDescType.Validators[0].(func(string) error)
}
