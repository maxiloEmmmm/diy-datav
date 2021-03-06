// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/assets"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/predicate"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
)

// AssetsUpdate is the builder for updating Assets entities.
type AssetsUpdate struct {
	config
	hooks    []Hook
	mutation *AssetsMutation
}

// Where adds a new predicate for the AssetsUpdate builder.
func (au *AssetsUpdate) Where(ps ...predicate.Assets) *AssetsUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetPath sets the "path" field.
func (au *AssetsUpdate) SetPath(s string) *AssetsUpdate {
	au.mutation.SetPath(s)
	return au
}

// SetExt sets the "ext" field.
func (au *AssetsUpdate) SetExt(s string) *AssetsUpdate {
	au.mutation.SetExt(s)
	return au
}

// SetType sets the "type" field.
func (au *AssetsUpdate) SetType(s string) *AssetsUpdate {
	au.mutation.SetType(s)
	return au
}

// AddViewIDs adds the "view" edge to the View entity by IDs.
func (au *AssetsUpdate) AddViewIDs(ids ...int) *AssetsUpdate {
	au.mutation.AddViewIDs(ids...)
	return au
}

// AddView adds the "view" edges to the View entity.
func (au *AssetsUpdate) AddView(v ...*View) *AssetsUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.AddViewIDs(ids...)
}

// Mutation returns the AssetsMutation object of the builder.
func (au *AssetsUpdate) Mutation() *AssetsMutation {
	return au.mutation
}

// ClearView clears all "view" edges to the View entity.
func (au *AssetsUpdate) ClearView() *AssetsUpdate {
	au.mutation.ClearView()
	return au
}

// RemoveViewIDs removes the "view" edge to View entities by IDs.
func (au *AssetsUpdate) RemoveViewIDs(ids ...int) *AssetsUpdate {
	au.mutation.RemoveViewIDs(ids...)
	return au
}

// RemoveView removes "view" edges to View entities.
func (au *AssetsUpdate) RemoveView(v ...*View) *AssetsUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.RemoveViewIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AssetsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AssetsUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AssetsUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AssetsUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AssetsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   assets.Table,
			Columns: assets.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: assets.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assets.FieldPath,
		})
	}
	if value, ok := au.mutation.Ext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assets.FieldExt,
		})
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assets.FieldType,
		})
	}
	if au.mutation.ViewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   assets.ViewTable,
			Columns: []string{assets.ViewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: view.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedViewIDs(); len(nodes) > 0 && !au.mutation.ViewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   assets.ViewTable,
			Columns: []string{assets.ViewColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ViewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   assets.ViewTable,
			Columns: []string{assets.ViewColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assets.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AssetsUpdateOne is the builder for updating a single Assets entity.
type AssetsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetsMutation
}

// SetPath sets the "path" field.
func (auo *AssetsUpdateOne) SetPath(s string) *AssetsUpdateOne {
	auo.mutation.SetPath(s)
	return auo
}

// SetExt sets the "ext" field.
func (auo *AssetsUpdateOne) SetExt(s string) *AssetsUpdateOne {
	auo.mutation.SetExt(s)
	return auo
}

// SetType sets the "type" field.
func (auo *AssetsUpdateOne) SetType(s string) *AssetsUpdateOne {
	auo.mutation.SetType(s)
	return auo
}

// AddViewIDs adds the "view" edge to the View entity by IDs.
func (auo *AssetsUpdateOne) AddViewIDs(ids ...int) *AssetsUpdateOne {
	auo.mutation.AddViewIDs(ids...)
	return auo
}

// AddView adds the "view" edges to the View entity.
func (auo *AssetsUpdateOne) AddView(v ...*View) *AssetsUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.AddViewIDs(ids...)
}

// Mutation returns the AssetsMutation object of the builder.
func (auo *AssetsUpdateOne) Mutation() *AssetsMutation {
	return auo.mutation
}

// ClearView clears all "view" edges to the View entity.
func (auo *AssetsUpdateOne) ClearView() *AssetsUpdateOne {
	auo.mutation.ClearView()
	return auo
}

// RemoveViewIDs removes the "view" edge to View entities by IDs.
func (auo *AssetsUpdateOne) RemoveViewIDs(ids ...int) *AssetsUpdateOne {
	auo.mutation.RemoveViewIDs(ids...)
	return auo
}

// RemoveView removes "view" edges to View entities.
func (auo *AssetsUpdateOne) RemoveView(v ...*View) *AssetsUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.RemoveViewIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AssetsUpdateOne) Select(field string, fields ...string) *AssetsUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Assets entity.
func (auo *AssetsUpdateOne) Save(ctx context.Context) (*Assets, error) {
	var (
		err  error
		node *Assets
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AssetsUpdateOne) SaveX(ctx context.Context) *Assets {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AssetsUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AssetsUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AssetsUpdateOne) sqlSave(ctx context.Context) (_node *Assets, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   assets.Table,
			Columns: assets.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: assets.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Assets.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assets.FieldID)
		for _, f := range fields {
			if !assets.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != assets.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assets.FieldPath,
		})
	}
	if value, ok := auo.mutation.Ext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assets.FieldExt,
		})
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assets.FieldType,
		})
	}
	if auo.mutation.ViewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   assets.ViewTable,
			Columns: []string{assets.ViewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: view.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedViewIDs(); len(nodes) > 0 && !auo.mutation.ViewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   assets.ViewTable,
			Columns: []string{assets.ViewColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ViewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   assets.ViewTable,
			Columns: []string{assets.ViewColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Assets{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assets.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
