// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/predicate"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
)

// ViewBlockDelete is the builder for deleting a ViewBlock entity.
type ViewBlockDelete struct {
	config
	hooks    []Hook
	mutation *ViewBlockMutation
}

// Where adds a new predicate to the ViewBlockDelete builder.
func (vbd *ViewBlockDelete) Where(ps ...predicate.ViewBlock) *ViewBlockDelete {
	vbd.mutation.predicates = append(vbd.mutation.predicates, ps...)
	return vbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vbd *ViewBlockDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vbd.hooks) == 0 {
		affected, err = vbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ViewBlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vbd.mutation = mutation
			affected, err = vbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vbd.hooks) - 1; i >= 0; i-- {
			mut = vbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vbd *ViewBlockDelete) ExecX(ctx context.Context) int {
	n, err := vbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vbd *ViewBlockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: viewblock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: viewblock.FieldID,
			},
		},
	}
	if ps := vbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vbd.driver, _spec)
}

// ViewBlockDeleteOne is the builder for deleting a single ViewBlock entity.
type ViewBlockDeleteOne struct {
	vbd *ViewBlockDelete
}

// Exec executes the deletion query.
func (vbdo *ViewBlockDeleteOne) Exec(ctx context.Context) error {
	n, err := vbdo.vbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{viewblock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vbdo *ViewBlockDeleteOne) ExecX(ctx context.Context) {
	vbdo.vbd.ExecX(ctx)
}
