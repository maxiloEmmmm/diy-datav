// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssetsColumns holds the columns for the "assets" table.
	AssetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "path", Type: field.TypeString},
		{Name: "ext", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:        "assets",
		Columns:     AssetsColumns,
		PrimaryKey:  []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DataSetsColumns holds the columns for the "data_sets" table.
	DataSetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Size: 255},
		{Name: "title", Type: field.TypeString, Size: 2147483647},
		{Name: "config", Type: field.TypeString, Size: 2147483647},
		{Name: "view_block_dataset", Type: field.TypeInt, Nullable: true},
	}
	// DataSetsTable holds the schema information for the "data_sets" table.
	DataSetsTable = &schema.Table{
		Name:       "data_sets",
		Columns:    DataSetsColumns,
		PrimaryKey: []*schema.Column{DataSetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "data_sets_view_blocks_dataset",
				Columns:    []*schema.Column{DataSetsColumns[4]},
				RefColumns: []*schema.Column{ViewBlocksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TypeConfigsColumns holds the columns for the "type_configs" table.
	TypeConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Size: 255},
		{Name: "config", Type: field.TypeString, Size: 2147483647},
	}
	// TypeConfigsTable holds the schema information for the "type_configs" table.
	TypeConfigsTable = &schema.Table{
		Name:        "type_configs",
		Columns:     TypeConfigsColumns,
		PrimaryKey:  []*schema.Column{TypeConfigsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ViewsColumns holds the columns for the "views" table.
	ViewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "desc", Type: field.TypeString, Size: 255},
		{Name: "config", Type: field.TypeString, Size: 2147483647},
		{Name: "assets_view", Type: field.TypeInt, Nullable: true},
	}
	// ViewsTable holds the schema information for the "views" table.
	ViewsTable = &schema.Table{
		Name:       "views",
		Columns:    ViewsColumns,
		PrimaryKey: []*schema.Column{ViewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "views_assets_view",
				Columns:    []*schema.Column{ViewsColumns[3]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ViewBlocksColumns holds the columns for the "view_blocks" table.
	ViewBlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Size: 255},
		{Name: "config", Type: field.TypeString, Size: 2147483647},
		{Name: "view_blocks", Type: field.TypeInt, Nullable: true},
	}
	// ViewBlocksTable holds the schema information for the "view_blocks" table.
	ViewBlocksTable = &schema.Table{
		Name:       "view_blocks",
		Columns:    ViewBlocksColumns,
		PrimaryKey: []*schema.Column{ViewBlocksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "view_blocks_views_blocks",
				Columns:    []*schema.Column{ViewBlocksColumns[3]},
				RefColumns: []*schema.Column{ViewsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssetsTable,
		DataSetsTable,
		TypeConfigsTable,
		ViewsTable,
		ViewBlocksTable,
	}
)

func init() {
	DataSetsTable.ForeignKeys[0].RefTable = ViewBlocksTable
	ViewsTable.ForeignKeys[0].RefTable = AssetsTable
	ViewBlocksTable.ForeignKeys[0].RefTable = ViewsTable
}
