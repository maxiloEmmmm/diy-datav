// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/maxiloEmmmm/diy-datav/pkg/model"
)

// The AssetsFunc type is an adapter to allow the use of ordinary
// function as Assets mutator.
type AssetsFunc func(context.Context, *model.AssetsMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f AssetsFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.AssetsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.AssetsMutation", m)
	}
	return f(ctx, mv)
}

// The DataSetFunc type is an adapter to allow the use of ordinary
// function as DataSet mutator.
type DataSetFunc func(context.Context, *model.DataSetMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f DataSetFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.DataSetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.DataSetMutation", m)
	}
	return f(ctx, mv)
}

// The TypeConfigFunc type is an adapter to allow the use of ordinary
// function as TypeConfig mutator.
type TypeConfigFunc func(context.Context, *model.TypeConfigMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f TypeConfigFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.TypeConfigMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.TypeConfigMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *model.UserMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.UserMutation", m)
	}
	return f(ctx, mv)
}

// The ViewFunc type is an adapter to allow the use of ordinary
// function as View mutator.
type ViewFunc func(context.Context, *model.ViewMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ViewFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.ViewMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ViewMutation", m)
	}
	return f(ctx, mv)
}

// The ViewBlockFunc type is an adapter to allow the use of ordinary
// function as ViewBlock mutator.
type ViewBlockFunc func(context.Context, *model.ViewBlockMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ViewBlockFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.ViewBlockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ViewBlockMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, model.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op model.Op) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk model.Hook, cond Condition) model.Hook {
	return func(next model.Mutator) model.Mutator {
		return model.MutateFunc(func(ctx context.Context, m model.Mutation) (model.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, model.Delete|model.Create)
//
func On(hk model.Hook, op model.Op) model.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, model.Update|model.UpdateOne)
//
func Unless(hk model.Hook, op model.Op) model.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) model.Hook {
	return func(model.Mutator) model.Mutator {
		return model.MutateFunc(func(context.Context, model.Mutation) (model.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []model.Hook {
//		return []model.Hook{
//			Reject(model.Delete|model.Update),
//		}
//	}
//
func Reject(op model.Op) model.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []model.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...model.Hook) Chain {
	return Chain{append([]model.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() model.Hook {
	return func(mutator model.Mutator) model.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...model.Hook) Chain {
	newHooks := make([]model.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
