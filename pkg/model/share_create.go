// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/share"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/user"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
)

// ShareCreate is the builder for creating a Share entity.
type ShareCreate struct {
	config
	mutation *ShareMutation
	hooks    []Hook
}

// SetEndAt sets the "end_at" field.
func (sc *ShareCreate) SetEndAt(t time.Time) *ShareCreate {
	sc.mutation.SetEndAt(t)
	return sc
}

// SetViewID sets the "view" edge to the View entity by ID.
func (sc *ShareCreate) SetViewID(id int) *ShareCreate {
	sc.mutation.SetViewID(id)
	return sc
}

// SetNillableViewID sets the "view" edge to the View entity by ID if the given value is not nil.
func (sc *ShareCreate) SetNillableViewID(id *int) *ShareCreate {
	if id != nil {
		sc = sc.SetViewID(*id)
	}
	return sc
}

// SetView sets the "view" edge to the View entity.
func (sc *ShareCreate) SetView(v *View) *ShareCreate {
	return sc.SetViewID(v.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (sc *ShareCreate) SetCreatorID(id int) *ShareCreate {
	sc.mutation.SetCreatorID(id)
	return sc
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (sc *ShareCreate) SetNillableCreatorID(id *int) *ShareCreate {
	if id != nil {
		sc = sc.SetCreatorID(*id)
	}
	return sc
}

// SetCreator sets the "creator" edge to the User entity.
func (sc *ShareCreate) SetCreator(u *User) *ShareCreate {
	return sc.SetCreatorID(u.ID)
}

// Mutation returns the ShareMutation object of the builder.
func (sc *ShareCreate) Mutation() *ShareMutation {
	return sc.mutation
}

// Save creates the Share in the database.
func (sc *ShareCreate) Save(ctx context.Context) (*Share, error) {
	var (
		err  error
		node *Share
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShareCreate) SaveX(ctx context.Context) *Share {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShareCreate) check() error {
	if _, ok := sc.mutation.EndAt(); !ok {
		return &ValidationError{Name: "end_at", err: errors.New("model: missing required field \"end_at\"")}
	}
	return nil
}

func (sc *ShareCreate) sqlSave(ctx context.Context) (*Share, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ShareCreate) createSpec() (*Share, *sqlgraph.CreateSpec) {
	var (
		_node = &Share{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: share.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: share.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.EndAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: share.FieldEndAt,
		})
		_node.EndAt = value
	}
	if nodes := sc.mutation.ViewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   share.ViewTable,
			Columns: []string{share.ViewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: view.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.view_share = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   share.CreatorTable,
			Columns: []string{share.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_share = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShareCreateBulk is the builder for creating many Share entities in bulk.
type ShareCreateBulk struct {
	config
	builders []*ShareCreate
}

// Save creates the Share entities in the database.
func (scb *ShareCreateBulk) Save(ctx context.Context) ([]*Share, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Share, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShareMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShareCreateBulk) SaveX(ctx context.Context) []*Share {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
