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
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/x/gidx"
)

// Representation of a status namespace. Status namespaces are used group status data that is provided by a resource provider.
type StatusNamespace struct {
	config `json:"-"`
	// ID of the ent.
	// The ID for the status namespace.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the status namespace.
	Name string `json:"name,omitempty"`
	// The ID for the tenant for this status namespace.
	ResourceProviderID gidx.PrefixedID `json:"resource_provider_id,omitempty"`
	// Flag for if this namespace is private.
	Private bool `json:"private,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatusNamespaceQuery when eager-loading is set.
	Edges        StatusNamespaceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StatusNamespaceEdges holds the relations/edges for other nodes in the graph.
type StatusNamespaceEdges struct {
	// Statuses holds the value of the statuses edge.
	Statuses []*Status `json:"statuses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool

	namedStatuses map[string][]*Status
}

// StatusesOrErr returns the Statuses value or an error if the edge
// was not loaded in eager-loading.
func (e StatusNamespaceEdges) StatusesOrErr() ([]*Status, error) {
	if e.loadedTypes[0] {
		return e.Statuses, nil
	}
	return nil, &NotLoadedError{edge: "statuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StatusNamespace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case statusnamespace.FieldID, statusnamespace.FieldResourceProviderID:
			values[i] = new(gidx.PrefixedID)
		case statusnamespace.FieldPrivate:
			values[i] = new(sql.NullBool)
		case statusnamespace.FieldName:
			values[i] = new(sql.NullString)
		case statusnamespace.FieldCreatedAt, statusnamespace.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StatusNamespace fields.
func (sn *StatusNamespace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case statusnamespace.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sn.ID = *value
			}
		case statusnamespace.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sn.CreatedAt = value.Time
			}
		case statusnamespace.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sn.UpdatedAt = value.Time
			}
		case statusnamespace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sn.Name = value.String
			}
		case statusnamespace.FieldResourceProviderID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field resource_provider_id", values[i])
			} else if value != nil {
				sn.ResourceProviderID = *value
			}
		case statusnamespace.FieldPrivate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field private", values[i])
			} else if value.Valid {
				sn.Private = value.Bool
			}
		default:
			sn.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StatusNamespace.
// This includes values selected through modifiers, order, etc.
func (sn *StatusNamespace) Value(name string) (ent.Value, error) {
	return sn.selectValues.Get(name)
}

// QueryStatuses queries the "statuses" edge of the StatusNamespace entity.
func (sn *StatusNamespace) QueryStatuses() *StatusQuery {
	return NewStatusNamespaceClient(sn.config).QueryStatuses(sn)
}

// Update returns a builder for updating this StatusNamespace.
// Note that you need to call StatusNamespace.Unwrap() before calling this method if this StatusNamespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (sn *StatusNamespace) Update() *StatusNamespaceUpdateOne {
	return NewStatusNamespaceClient(sn.config).UpdateOne(sn)
}

// Unwrap unwraps the StatusNamespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sn *StatusNamespace) Unwrap() *StatusNamespace {
	_tx, ok := sn.config.driver.(*txDriver)
	if !ok {
		panic("generated: StatusNamespace is not a transactional entity")
	}
	sn.config.driver = _tx.drv
	return sn
}

// String implements the fmt.Stringer.
func (sn *StatusNamespace) String() string {
	var builder strings.Builder
	builder.WriteString("StatusNamespace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sn.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sn.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sn.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sn.Name)
	builder.WriteString(", ")
	builder.WriteString("resource_provider_id=")
	builder.WriteString(fmt.Sprintf("%v", sn.ResourceProviderID))
	builder.WriteString(", ")
	builder.WriteString("private=")
	builder.WriteString(fmt.Sprintf("%v", sn.Private))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (sn StatusNamespace) IsEntity() {}

// NamedStatuses returns the Statuses named value or an error if the edge was not
// loaded in eager-loading with this name.
func (sn *StatusNamespace) NamedStatuses(name string) ([]*Status, error) {
	if sn.Edges.namedStatuses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := sn.Edges.namedStatuses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (sn *StatusNamespace) appendNamedStatuses(name string, edges ...*Status) {
	if sn.Edges.namedStatuses == nil {
		sn.Edges.namedStatuses = make(map[string][]*Status)
	}
	if len(edges) == 0 {
		sn.Edges.namedStatuses[name] = []*Status{}
	} else {
		sn.Edges.namedStatuses[name] = append(sn.Edges.namedStatuses[name], edges...)
	}
}

// StatusNamespaces is a parsable slice of StatusNamespace.
type StatusNamespaces []*StatusNamespace