// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"web-stash-api/ent/bag"
	"web-stash-api/ent/bagitem"
	"web-stash-api/ent/subitem"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BagItemCreate is the builder for creating a BagItem entity.
type BagItemCreate struct {
	config
	mutation *BagItemMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bic *BagItemCreate) SetName(s string) *BagItemCreate {
	bic.mutation.SetName(s)
	return bic
}

// SetDescription sets the "description" field.
func (bic *BagItemCreate) SetDescription(s string) *BagItemCreate {
	bic.mutation.SetDescription(s)
	return bic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bic *BagItemCreate) SetNillableDescription(s *string) *BagItemCreate {
	if s != nil {
		bic.SetDescription(*s)
	}
	return bic
}

// SetIcon sets the "icon" field.
func (bic *BagItemCreate) SetIcon(s string) *BagItemCreate {
	bic.mutation.SetIcon(s)
	return bic
}

// SetLink sets the "link" field.
func (bic *BagItemCreate) SetLink(s string) *BagItemCreate {
	bic.mutation.SetLink(s)
	return bic
}

// SetImage sets the "image" field.
func (bic *BagItemCreate) SetImage(s string) *BagItemCreate {
	bic.mutation.SetImage(s)
	return bic
}

// SetID sets the "id" field.
func (bic *BagItemCreate) SetID(u uuid.UUID) *BagItemCreate {
	bic.mutation.SetID(u)
	return bic
}

// SetBagID sets the "bag" edge to the Bag entity by ID.
func (bic *BagItemCreate) SetBagID(id uuid.UUID) *BagItemCreate {
	bic.mutation.SetBagID(id)
	return bic
}

// SetNillableBagID sets the "bag" edge to the Bag entity by ID if the given value is not nil.
func (bic *BagItemCreate) SetNillableBagID(id *uuid.UUID) *BagItemCreate {
	if id != nil {
		bic = bic.SetBagID(*id)
	}
	return bic
}

// SetBag sets the "bag" edge to the Bag entity.
func (bic *BagItemCreate) SetBag(b *Bag) *BagItemCreate {
	return bic.SetBagID(b.ID)
}

// AddSubItemIDs adds the "sub_items" edge to the SubItem entity by IDs.
func (bic *BagItemCreate) AddSubItemIDs(ids ...uuid.UUID) *BagItemCreate {
	bic.mutation.AddSubItemIDs(ids...)
	return bic
}

// AddSubItems adds the "sub_items" edges to the SubItem entity.
func (bic *BagItemCreate) AddSubItems(s ...*SubItem) *BagItemCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bic.AddSubItemIDs(ids...)
}

// Mutation returns the BagItemMutation object of the builder.
func (bic *BagItemCreate) Mutation() *BagItemMutation {
	return bic.mutation
}

// Save creates the BagItem in the database.
func (bic *BagItemCreate) Save(ctx context.Context) (*BagItem, error) {
	var (
		err  error
		node *BagItem
	)
	bic.defaults()
	if len(bic.hooks) == 0 {
		if err = bic.check(); err != nil {
			return nil, err
		}
		node, err = bic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BagItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bic.check(); err != nil {
				return nil, err
			}
			bic.mutation = mutation
			node, err = bic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bic.hooks) - 1; i >= 0; i-- {
			mut = bic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bic *BagItemCreate) SaveX(ctx context.Context) *BagItem {
	v, err := bic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bic *BagItemCreate) defaults() {
	if _, ok := bic.mutation.ID(); !ok {
		v := bagitem.DefaultID()
		bic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bic *BagItemCreate) check() error {
	if _, ok := bic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := bic.mutation.Name(); ok {
		if err := bagitem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := bic.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New("ent: missing required field \"icon\"")}
	}
	if v, ok := bic.mutation.Icon(); ok {
		if err := bagitem.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf("ent: validator failed for field \"icon\": %w", err)}
		}
	}
	if _, ok := bic.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New("ent: missing required field \"link\"")}
	}
	if v, ok := bic.mutation.Link(); ok {
		if err := bagitem.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf("ent: validator failed for field \"link\": %w", err)}
		}
	}
	if _, ok := bic.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New("ent: missing required field \"image\"")}
	}
	if v, ok := bic.mutation.Image(); ok {
		if err := bagitem.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf("ent: validator failed for field \"image\": %w", err)}
		}
	}
	return nil
}

func (bic *BagItemCreate) sqlSave(ctx context.Context) (*BagItem, error) {
	_node, _spec := bic.createSpec()
	if err := sqlgraph.CreateNode(ctx, bic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (bic *BagItemCreate) createSpec() (*BagItem, *sqlgraph.CreateSpec) {
	var (
		_node = &BagItem{config: bic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bagitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: bagitem.FieldID,
			},
		}
	)
	if id, ok := bic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bagitem.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bic.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bagitem.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := bic.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bagitem.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := bic.mutation.Link(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bagitem.FieldLink,
		})
		_node.Link = value
	}
	if value, ok := bic.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bagitem.FieldImage,
		})
		_node.Image = value
	}
	if nodes := bic.mutation.BagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bagitem.BagTable,
			Columns: []string{bagitem.BagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bic.mutation.SubItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bagitem.SubItemsTable,
			Columns: []string{bagitem.SubItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BagItemCreateBulk is the builder for creating many BagItem entities in bulk.
type BagItemCreateBulk struct {
	config
	builders []*BagItemCreate
}

// Save creates the BagItem entities in the database.
func (bicb *BagItemCreateBulk) Save(ctx context.Context) ([]*BagItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bicb.builders))
	nodes := make([]*BagItem, len(bicb.builders))
	mutators := make([]Mutator, len(bicb.builders))
	for i := range bicb.builders {
		func(i int, root context.Context) {
			builder := bicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BagItemMutation)
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
					_, err = mutators[i+1].Mutate(root, bicb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bicb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, bicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bicb *BagItemCreateBulk) SaveX(ctx context.Context) []*BagItem {
	v, err := bicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}