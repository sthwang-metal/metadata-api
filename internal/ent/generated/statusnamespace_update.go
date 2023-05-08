// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/x/gidx"
)

// StatusNamespaceUpdate is the builder for updating StatusNamespace entities.
type StatusNamespaceUpdate struct {
	config
	hooks    []Hook
	mutation *StatusNamespaceMutation
}

// Where appends a list predicates to the StatusNamespaceUpdate builder.
func (snu *StatusNamespaceUpdate) Where(ps ...predicate.StatusNamespace) *StatusNamespaceUpdate {
	snu.mutation.Where(ps...)
	return snu
}

// SetName sets the "name" field.
func (snu *StatusNamespaceUpdate) SetName(s string) *StatusNamespaceUpdate {
	snu.mutation.SetName(s)
	return snu
}

// SetPrivate sets the "private" field.
func (snu *StatusNamespaceUpdate) SetPrivate(b bool) *StatusNamespaceUpdate {
	snu.mutation.SetPrivate(b)
	return snu
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (snu *StatusNamespaceUpdate) SetNillablePrivate(b *bool) *StatusNamespaceUpdate {
	if b != nil {
		snu.SetPrivate(*b)
	}
	return snu
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (snu *StatusNamespaceUpdate) AddStatusIDs(ids ...gidx.PrefixedID) *StatusNamespaceUpdate {
	snu.mutation.AddStatusIDs(ids...)
	return snu
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (snu *StatusNamespaceUpdate) AddStatuses(s ...*Status) *StatusNamespaceUpdate {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snu.AddStatusIDs(ids...)
}

// Mutation returns the StatusNamespaceMutation object of the builder.
func (snu *StatusNamespaceUpdate) Mutation() *StatusNamespaceMutation {
	return snu.mutation
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (snu *StatusNamespaceUpdate) ClearStatuses() *StatusNamespaceUpdate {
	snu.mutation.ClearStatuses()
	return snu
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (snu *StatusNamespaceUpdate) RemoveStatusIDs(ids ...gidx.PrefixedID) *StatusNamespaceUpdate {
	snu.mutation.RemoveStatusIDs(ids...)
	return snu
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (snu *StatusNamespaceUpdate) RemoveStatuses(s ...*Status) *StatusNamespaceUpdate {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snu.RemoveStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (snu *StatusNamespaceUpdate) Save(ctx context.Context) (int, error) {
	snu.defaults()
	return withHooks[int, StatusNamespaceMutation](ctx, snu.sqlSave, snu.mutation, snu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (snu *StatusNamespaceUpdate) SaveX(ctx context.Context) int {
	affected, err := snu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (snu *StatusNamespaceUpdate) Exec(ctx context.Context) error {
	_, err := snu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (snu *StatusNamespaceUpdate) ExecX(ctx context.Context) {
	if err := snu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (snu *StatusNamespaceUpdate) defaults() {
	if _, ok := snu.mutation.UpdatedAt(); !ok {
		v := statusnamespace.UpdateDefaultUpdatedAt()
		snu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (snu *StatusNamespaceUpdate) check() error {
	if v, ok := snu.mutation.Name(); ok {
		if err := statusnamespace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "StatusNamespace.name": %w`, err)}
		}
	}
	return nil
}

func (snu *StatusNamespaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := snu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(statusnamespace.Table, statusnamespace.Columns, sqlgraph.NewFieldSpec(statusnamespace.FieldID, field.TypeString))
	if ps := snu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := snu.mutation.UpdatedAt(); ok {
		_spec.SetField(statusnamespace.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := snu.mutation.Name(); ok {
		_spec.SetField(statusnamespace.FieldName, field.TypeString, value)
	}
	if value, ok := snu.mutation.Private(); ok {
		_spec.SetField(statusnamespace.FieldPrivate, field.TypeBool, value)
	}
	if snu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   statusnamespace.StatusesTable,
			Columns: []string{statusnamespace.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snu.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !snu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   statusnamespace.StatusesTable,
			Columns: []string{statusnamespace.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snu.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   statusnamespace.StatusesTable,
			Columns: []string{statusnamespace.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, snu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statusnamespace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	snu.mutation.done = true
	return n, nil
}

// StatusNamespaceUpdateOne is the builder for updating a single StatusNamespace entity.
type StatusNamespaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StatusNamespaceMutation
}

// SetName sets the "name" field.
func (snuo *StatusNamespaceUpdateOne) SetName(s string) *StatusNamespaceUpdateOne {
	snuo.mutation.SetName(s)
	return snuo
}

// SetPrivate sets the "private" field.
func (snuo *StatusNamespaceUpdateOne) SetPrivate(b bool) *StatusNamespaceUpdateOne {
	snuo.mutation.SetPrivate(b)
	return snuo
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (snuo *StatusNamespaceUpdateOne) SetNillablePrivate(b *bool) *StatusNamespaceUpdateOne {
	if b != nil {
		snuo.SetPrivate(*b)
	}
	return snuo
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (snuo *StatusNamespaceUpdateOne) AddStatusIDs(ids ...gidx.PrefixedID) *StatusNamespaceUpdateOne {
	snuo.mutation.AddStatusIDs(ids...)
	return snuo
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (snuo *StatusNamespaceUpdateOne) AddStatuses(s ...*Status) *StatusNamespaceUpdateOne {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snuo.AddStatusIDs(ids...)
}

// Mutation returns the StatusNamespaceMutation object of the builder.
func (snuo *StatusNamespaceUpdateOne) Mutation() *StatusNamespaceMutation {
	return snuo.mutation
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (snuo *StatusNamespaceUpdateOne) ClearStatuses() *StatusNamespaceUpdateOne {
	snuo.mutation.ClearStatuses()
	return snuo
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (snuo *StatusNamespaceUpdateOne) RemoveStatusIDs(ids ...gidx.PrefixedID) *StatusNamespaceUpdateOne {
	snuo.mutation.RemoveStatusIDs(ids...)
	return snuo
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (snuo *StatusNamespaceUpdateOne) RemoveStatuses(s ...*Status) *StatusNamespaceUpdateOne {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snuo.RemoveStatusIDs(ids...)
}

// Where appends a list predicates to the StatusNamespaceUpdate builder.
func (snuo *StatusNamespaceUpdateOne) Where(ps ...predicate.StatusNamespace) *StatusNamespaceUpdateOne {
	snuo.mutation.Where(ps...)
	return snuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (snuo *StatusNamespaceUpdateOne) Select(field string, fields ...string) *StatusNamespaceUpdateOne {
	snuo.fields = append([]string{field}, fields...)
	return snuo
}

// Save executes the query and returns the updated StatusNamespace entity.
func (snuo *StatusNamespaceUpdateOne) Save(ctx context.Context) (*StatusNamespace, error) {
	snuo.defaults()
	return withHooks[*StatusNamespace, StatusNamespaceMutation](ctx, snuo.sqlSave, snuo.mutation, snuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (snuo *StatusNamespaceUpdateOne) SaveX(ctx context.Context) *StatusNamespace {
	node, err := snuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (snuo *StatusNamespaceUpdateOne) Exec(ctx context.Context) error {
	_, err := snuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (snuo *StatusNamespaceUpdateOne) ExecX(ctx context.Context) {
	if err := snuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (snuo *StatusNamespaceUpdateOne) defaults() {
	if _, ok := snuo.mutation.UpdatedAt(); !ok {
		v := statusnamespace.UpdateDefaultUpdatedAt()
		snuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (snuo *StatusNamespaceUpdateOne) check() error {
	if v, ok := snuo.mutation.Name(); ok {
		if err := statusnamespace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "StatusNamespace.name": %w`, err)}
		}
	}
	return nil
}

func (snuo *StatusNamespaceUpdateOne) sqlSave(ctx context.Context) (_node *StatusNamespace, err error) {
	if err := snuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(statusnamespace.Table, statusnamespace.Columns, sqlgraph.NewFieldSpec(statusnamespace.FieldID, field.TypeString))
	id, ok := snuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "StatusNamespace.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := snuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statusnamespace.FieldID)
		for _, f := range fields {
			if !statusnamespace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != statusnamespace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := snuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := snuo.mutation.UpdatedAt(); ok {
		_spec.SetField(statusnamespace.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := snuo.mutation.Name(); ok {
		_spec.SetField(statusnamespace.FieldName, field.TypeString, value)
	}
	if value, ok := snuo.mutation.Private(); ok {
		_spec.SetField(statusnamespace.FieldPrivate, field.TypeBool, value)
	}
	if snuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   statusnamespace.StatusesTable,
			Columns: []string{statusnamespace.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snuo.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !snuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   statusnamespace.StatusesTable,
			Columns: []string{statusnamespace.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snuo.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   statusnamespace.StatusesTable,
			Columns: []string{statusnamespace.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StatusNamespace{config: snuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, snuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statusnamespace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	snuo.mutation.done = true
	return _node, nil
}