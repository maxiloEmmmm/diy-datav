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
	"github.com/maxiloEmmmm/diy-datav/pkg/model/share"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
)

// ViewUpdate is the builder for updating View entities.
type ViewUpdate struct {
	config
	hooks    []Hook
	mutation *ViewMutation
}

// Where adds a new predicate for the ViewUpdate builder.
func (vu *ViewUpdate) Where(ps ...predicate.View) *ViewUpdate {
	vu.mutation.predicates = append(vu.mutation.predicates, ps...)
	return vu
}

// SetDesc sets the "desc" field.
func (vu *ViewUpdate) SetDesc(s string) *ViewUpdate {
	vu.mutation.SetDesc(s)
	return vu
}

// SetConfig sets the "config" field.
func (vu *ViewUpdate) SetConfig(s string) *ViewUpdate {
	vu.mutation.SetConfig(s)
	return vu
}

// SetBgID sets the "bg" edge to the Assets entity by ID.
func (vu *ViewUpdate) SetBgID(id int) *ViewUpdate {
	vu.mutation.SetBgID(id)
	return vu
}

// SetNillableBgID sets the "bg" edge to the Assets entity by ID if the given value is not nil.
func (vu *ViewUpdate) SetNillableBgID(id *int) *ViewUpdate {
	if id != nil {
		vu = vu.SetBgID(*id)
	}
	return vu
}

// SetBg sets the "bg" edge to the Assets entity.
func (vu *ViewUpdate) SetBg(a *Assets) *ViewUpdate {
	return vu.SetBgID(a.ID)
}

// AddBlockIDs adds the "blocks" edge to the ViewBlock entity by IDs.
func (vu *ViewUpdate) AddBlockIDs(ids ...int) *ViewUpdate {
	vu.mutation.AddBlockIDs(ids...)
	return vu
}

// AddBlocks adds the "blocks" edges to the ViewBlock entity.
func (vu *ViewUpdate) AddBlocks(v ...*ViewBlock) *ViewUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.AddBlockIDs(ids...)
}

// AddShareIDs adds the "share" edge to the Share entity by IDs.
func (vu *ViewUpdate) AddShareIDs(ids ...int) *ViewUpdate {
	vu.mutation.AddShareIDs(ids...)
	return vu
}

// AddShare adds the "share" edges to the Share entity.
func (vu *ViewUpdate) AddShare(s ...*Share) *ViewUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return vu.AddShareIDs(ids...)
}

// Mutation returns the ViewMutation object of the builder.
func (vu *ViewUpdate) Mutation() *ViewMutation {
	return vu.mutation
}

// ClearBg clears the "bg" edge to the Assets entity.
func (vu *ViewUpdate) ClearBg() *ViewUpdate {
	vu.mutation.ClearBg()
	return vu
}

// ClearBlocks clears all "blocks" edges to the ViewBlock entity.
func (vu *ViewUpdate) ClearBlocks() *ViewUpdate {
	vu.mutation.ClearBlocks()
	return vu
}

// RemoveBlockIDs removes the "blocks" edge to ViewBlock entities by IDs.
func (vu *ViewUpdate) RemoveBlockIDs(ids ...int) *ViewUpdate {
	vu.mutation.RemoveBlockIDs(ids...)
	return vu
}

// RemoveBlocks removes "blocks" edges to ViewBlock entities.
func (vu *ViewUpdate) RemoveBlocks(v ...*ViewBlock) *ViewUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.RemoveBlockIDs(ids...)
}

// ClearShare clears all "share" edges to the Share entity.
func (vu *ViewUpdate) ClearShare() *ViewUpdate {
	vu.mutation.ClearShare()
	return vu
}

// RemoveShareIDs removes the "share" edge to Share entities by IDs.
func (vu *ViewUpdate) RemoveShareIDs(ids ...int) *ViewUpdate {
	vu.mutation.RemoveShareIDs(ids...)
	return vu
}

// RemoveShare removes "share" edges to Share entities.
func (vu *ViewUpdate) RemoveShare(s ...*Share) *ViewUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return vu.RemoveShareIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *ViewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		if err = vu.check(); err != nil {
			return 0, err
		}
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ViewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vu.check(); err != nil {
				return 0, err
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *ViewUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *ViewUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *ViewUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *ViewUpdate) check() error {
	if v, ok := vu.mutation.Desc(); ok {
		if err := view.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf("model: validator failed for field \"desc\": %w", err)}
		}
	}
	return nil
}

func (vu *ViewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   view.Table,
			Columns: view.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: view.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: view.FieldDesc,
		})
	}
	if value, ok := vu.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: view.FieldConfig,
		})
	}
	if vu.mutation.BgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.BgTable,
			Columns: []string{view.BgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: assets.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.BgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.BgTable,
			Columns: []string{view.BgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: assets.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.BlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.BlocksTable,
			Columns: []string{view.BlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: viewblock.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedBlocksIDs(); len(nodes) > 0 && !vu.mutation.BlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.BlocksTable,
			Columns: []string{view.BlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: viewblock.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.BlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.BlocksTable,
			Columns: []string{view.BlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: viewblock.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.ShareCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.ShareTable,
			Columns: []string{view.ShareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: share.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedShareIDs(); len(nodes) > 0 && !vu.mutation.ShareCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.ShareTable,
			Columns: []string{view.ShareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: share.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.ShareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.ShareTable,
			Columns: []string{view.ShareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: share.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{view.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ViewUpdateOne is the builder for updating a single View entity.
type ViewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ViewMutation
}

// SetDesc sets the "desc" field.
func (vuo *ViewUpdateOne) SetDesc(s string) *ViewUpdateOne {
	vuo.mutation.SetDesc(s)
	return vuo
}

// SetConfig sets the "config" field.
func (vuo *ViewUpdateOne) SetConfig(s string) *ViewUpdateOne {
	vuo.mutation.SetConfig(s)
	return vuo
}

// SetBgID sets the "bg" edge to the Assets entity by ID.
func (vuo *ViewUpdateOne) SetBgID(id int) *ViewUpdateOne {
	vuo.mutation.SetBgID(id)
	return vuo
}

// SetNillableBgID sets the "bg" edge to the Assets entity by ID if the given value is not nil.
func (vuo *ViewUpdateOne) SetNillableBgID(id *int) *ViewUpdateOne {
	if id != nil {
		vuo = vuo.SetBgID(*id)
	}
	return vuo
}

// SetBg sets the "bg" edge to the Assets entity.
func (vuo *ViewUpdateOne) SetBg(a *Assets) *ViewUpdateOne {
	return vuo.SetBgID(a.ID)
}

// AddBlockIDs adds the "blocks" edge to the ViewBlock entity by IDs.
func (vuo *ViewUpdateOne) AddBlockIDs(ids ...int) *ViewUpdateOne {
	vuo.mutation.AddBlockIDs(ids...)
	return vuo
}

// AddBlocks adds the "blocks" edges to the ViewBlock entity.
func (vuo *ViewUpdateOne) AddBlocks(v ...*ViewBlock) *ViewUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.AddBlockIDs(ids...)
}

// AddShareIDs adds the "share" edge to the Share entity by IDs.
func (vuo *ViewUpdateOne) AddShareIDs(ids ...int) *ViewUpdateOne {
	vuo.mutation.AddShareIDs(ids...)
	return vuo
}

// AddShare adds the "share" edges to the Share entity.
func (vuo *ViewUpdateOne) AddShare(s ...*Share) *ViewUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return vuo.AddShareIDs(ids...)
}

// Mutation returns the ViewMutation object of the builder.
func (vuo *ViewUpdateOne) Mutation() *ViewMutation {
	return vuo.mutation
}

// ClearBg clears the "bg" edge to the Assets entity.
func (vuo *ViewUpdateOne) ClearBg() *ViewUpdateOne {
	vuo.mutation.ClearBg()
	return vuo
}

// ClearBlocks clears all "blocks" edges to the ViewBlock entity.
func (vuo *ViewUpdateOne) ClearBlocks() *ViewUpdateOne {
	vuo.mutation.ClearBlocks()
	return vuo
}

// RemoveBlockIDs removes the "blocks" edge to ViewBlock entities by IDs.
func (vuo *ViewUpdateOne) RemoveBlockIDs(ids ...int) *ViewUpdateOne {
	vuo.mutation.RemoveBlockIDs(ids...)
	return vuo
}

// RemoveBlocks removes "blocks" edges to ViewBlock entities.
func (vuo *ViewUpdateOne) RemoveBlocks(v ...*ViewBlock) *ViewUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.RemoveBlockIDs(ids...)
}

// ClearShare clears all "share" edges to the Share entity.
func (vuo *ViewUpdateOne) ClearShare() *ViewUpdateOne {
	vuo.mutation.ClearShare()
	return vuo
}

// RemoveShareIDs removes the "share" edge to Share entities by IDs.
func (vuo *ViewUpdateOne) RemoveShareIDs(ids ...int) *ViewUpdateOne {
	vuo.mutation.RemoveShareIDs(ids...)
	return vuo
}

// RemoveShare removes "share" edges to Share entities.
func (vuo *ViewUpdateOne) RemoveShare(s ...*Share) *ViewUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return vuo.RemoveShareIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *ViewUpdateOne) Select(field string, fields ...string) *ViewUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated View entity.
func (vuo *ViewUpdateOne) Save(ctx context.Context) (*View, error) {
	var (
		err  error
		node *View
	)
	if len(vuo.hooks) == 0 {
		if err = vuo.check(); err != nil {
			return nil, err
		}
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ViewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuo.check(); err != nil {
				return nil, err
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *ViewUpdateOne) SaveX(ctx context.Context) *View {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *ViewUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *ViewUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *ViewUpdateOne) check() error {
	if v, ok := vuo.mutation.Desc(); ok {
		if err := view.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf("model: validator failed for field \"desc\": %w", err)}
		}
	}
	return nil
}

func (vuo *ViewUpdateOne) sqlSave(ctx context.Context) (_node *View, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   view.Table,
			Columns: view.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: view.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing View.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, view.FieldID)
		for _, f := range fields {
			if !view.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != view.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: view.FieldDesc,
		})
	}
	if value, ok := vuo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: view.FieldConfig,
		})
	}
	if vuo.mutation.BgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.BgTable,
			Columns: []string{view.BgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: assets.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.BgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.BgTable,
			Columns: []string{view.BgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: assets.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.BlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.BlocksTable,
			Columns: []string{view.BlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: viewblock.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedBlocksIDs(); len(nodes) > 0 && !vuo.mutation.BlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.BlocksTable,
			Columns: []string{view.BlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: viewblock.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.BlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.BlocksTable,
			Columns: []string{view.BlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: viewblock.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.ShareCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.ShareTable,
			Columns: []string{view.ShareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: share.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedShareIDs(); len(nodes) > 0 && !vuo.mutation.ShareCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.ShareTable,
			Columns: []string{view.ShareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: share.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.ShareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   view.ShareTable,
			Columns: []string{view.ShareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: share.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &View{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{view.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
