// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/schema"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
)

// TypeConfigCreate is the builder for creating a TypeConfig entity.
type TypeConfigCreate struct {
	config
	mutation *TypeConfigMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (tcc *TypeConfigCreate) SetType(s string) *TypeConfigCreate {
	tcc.mutation.SetType(s)
	return tcc
}

// SetConfig sets the "config" field.
func (tcc *TypeConfigCreate) SetConfig(s string) *TypeConfigCreate {
	tcc.mutation.SetConfig(s)
	return tcc
}

// SetID sets the "id" field.
func (tcc *TypeConfigCreate) SetID(sk schema.TypeKey) *TypeConfigCreate {
	tcc.mutation.SetID(sk)
	return tcc
}

// Mutation returns the TypeConfigMutation object of the builder.
func (tcc *TypeConfigCreate) Mutation() *TypeConfigMutation {
	return tcc.mutation
}

// Save creates the TypeConfig in the database.
func (tcc *TypeConfigCreate) Save(ctx context.Context) (*TypeConfig, error) {
	var (
		err  error
		node *TypeConfig
	)
	if len(tcc.hooks) == 0 {
		if err = tcc.check(); err != nil {
			return nil, err
		}
		node, err = tcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypeConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tcc.check(); err != nil {
				return nil, err
			}
			tcc.mutation = mutation
			node, err = tcc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tcc.hooks) - 1; i >= 0; i-- {
			mut = tcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tcc *TypeConfigCreate) SaveX(ctx context.Context) *TypeConfig {
	v, err := tcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (tcc *TypeConfigCreate) check() error {
	if _, ok := tcc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("model: missing required field \"type\"")}
	}
	if v, ok := tcc.mutation.GetType(); ok {
		if err := typeconfig.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("model: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := tcc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New("model: missing required field \"config\"")}
	}
	return nil
}

func (tcc *TypeConfigCreate) sqlSave(ctx context.Context) (*TypeConfig, error) {
	_node, _spec := tcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = schema.TypeKey(id)
	}
	return _node, nil
}

func (tcc *TypeConfigCreate) createSpec() (*TypeConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &TypeConfig{config: tcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: typeconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typeconfig.FieldID,
			},
		}
	)
	if id, ok := tcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tcc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typeconfig.FieldType,
		})
		_node.Type = value
	}
	if value, ok := tcc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typeconfig.FieldConfig,
		})
		_node.Config = value
	}
	return _node, _spec
}

// TypeConfigCreateBulk is the builder for creating many TypeConfig entities in bulk.
type TypeConfigCreateBulk struct {
	config
	builders []*TypeConfigCreate
}

// Save creates the TypeConfig entities in the database.
func (tccb *TypeConfigCreateBulk) Save(ctx context.Context) ([]*TypeConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tccb.builders))
	nodes := make([]*TypeConfig, len(tccb.builders))
	mutators := make([]Mutator, len(tccb.builders))
	for i := range tccb.builders {
		func(i int, root context.Context) {
			builder := tccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TypeConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = schema.TypeKey(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tccb *TypeConfigCreateBulk) SaveX(ctx context.Context) []*TypeConfig {
	v, err := tccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
