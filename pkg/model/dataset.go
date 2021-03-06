// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/dataset"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
)

// DataSet is the model entity for the DataSet schema.
type DataSet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Config holds the value of the "config" field.
	Config string `json:"config,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DataSetQuery when eager-loading is set.
	Edges              DataSetEdges `json:"edges"`
	view_block_dataset *int
}

// DataSetEdges holds the relations/edges for other nodes in the graph.
type DataSetEdges struct {
	// Block holds the value of the block edge.
	Block *ViewBlock `json:"block,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BlockOrErr returns the Block value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DataSetEdges) BlockOrErr() (*ViewBlock, error) {
	if e.loadedTypes[0] {
		if e.Block == nil {
			// The edge block was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: viewblock.Label}
		}
		return e.Block, nil
	}
	return nil, &NotLoadedError{edge: "block"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DataSet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dataset.FieldID:
			values[i] = new(sql.NullInt64)
		case dataset.FieldType, dataset.FieldTitle, dataset.FieldConfig:
			values[i] = new(sql.NullString)
		case dataset.ForeignKeys[0]: // view_block_dataset
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DataSet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DataSet fields.
func (ds *DataSet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dataset.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ds.ID = int(value.Int64)
		case dataset.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ds.Type = value.String
			}
		case dataset.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ds.Title = value.String
			}
		case dataset.FieldConfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value.Valid {
				ds.Config = value.String
			}
		case dataset.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field view_block_dataset", value)
			} else if value.Valid {
				ds.view_block_dataset = new(int)
				*ds.view_block_dataset = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBlock queries the "block" edge of the DataSet entity.
func (ds *DataSet) QueryBlock() *ViewBlockQuery {
	return (&DataSetClient{config: ds.config}).QueryBlock(ds)
}

// Update returns a builder for updating this DataSet.
// Note that you need to call DataSet.Unwrap() before calling this method if this DataSet
// was returned from a transaction, and the transaction was committed or rolled back.
func (ds *DataSet) Update() *DataSetUpdateOne {
	return (&DataSetClient{config: ds.config}).UpdateOne(ds)
}

// Unwrap unwraps the DataSet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ds *DataSet) Unwrap() *DataSet {
	tx, ok := ds.config.driver.(*txDriver)
	if !ok {
		panic("model: DataSet is not a transactional entity")
	}
	ds.config.driver = tx.drv
	return ds
}

// String implements the fmt.Stringer.
func (ds *DataSet) String() string {
	var builder strings.Builder
	builder.WriteString("DataSet(")
	builder.WriteString(fmt.Sprintf("id=%v", ds.ID))
	builder.WriteString(", type=")
	builder.WriteString(ds.Type)
	builder.WriteString(", title=")
	builder.WriteString(ds.Title)
	builder.WriteString(", config=")
	builder.WriteString(ds.Config)
	builder.WriteByte(')')
	return builder.String()
}

// DataSets is a parsable slice of DataSet.
type DataSets []*DataSet

func (ds DataSets) config(cfg config) {
	for _i := range ds {
		ds[_i].config = cfg
	}
}
