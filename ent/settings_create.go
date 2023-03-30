// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ppaanngggg/MagicConch/ent/settings"
)

// SettingsCreate is the builder for creating a Settings entity.
type SettingsCreate struct {
	config
	mutation *SettingsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetKey sets the "key" field.
func (sc *SettingsCreate) SetKey(s string) *SettingsCreate {
	sc.mutation.SetKey(s)
	return sc
}

// SetValues sets the "values" field.
func (sc *SettingsCreate) SetValues(s []string) *SettingsCreate {
	sc.mutation.SetValues(s)
	return sc
}

// Mutation returns the SettingsMutation object of the builder.
func (sc *SettingsCreate) Mutation() *SettingsMutation {
	return sc.mutation
}

// Save creates the Settings in the database.
func (sc *SettingsCreate) Save(ctx context.Context) (*Settings, error) {
	return withHooks[*Settings, SettingsMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingsCreate) SaveX(ctx context.Context) *Settings {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingsCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingsCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingsCreate) check() error {
	if _, ok := sc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Settings.key"`)}
	}
	if _, ok := sc.mutation.Values(); !ok {
		return &ValidationError{Name: "values", err: errors.New(`ent: missing required field "Settings.values"`)}
	}
	return nil
}

func (sc *SettingsCreate) sqlSave(ctx context.Context) (*Settings, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingsCreate) createSpec() (*Settings, *sqlgraph.CreateSpec) {
	var (
		_node = &Settings{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(settings.Table, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Key(); ok {
		_spec.SetField(settings.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := sc.mutation.Values(); ok {
		_spec.SetField(settings.FieldValues, field.TypeJSON, value)
		_node.Values = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Settings.Create().
//		SetKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingsUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
func (sc *SettingsCreate) OnConflict(opts ...sql.ConflictOption) *SettingsUpsertOne {
	sc.conflict = opts
	return &SettingsUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SettingsCreate) OnConflictColumns(columns ...string) *SettingsUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SettingsUpsertOne{
		create: sc,
	}
}

type (
	// SettingsUpsertOne is the builder for "upsert"-ing
	//  one Settings node.
	SettingsUpsertOne struct {
		create *SettingsCreate
	}

	// SettingsUpsert is the "OnConflict" setter.
	SettingsUpsert struct {
		*sql.UpdateSet
	}
)

// SetKey sets the "key" field.
func (u *SettingsUpsert) SetKey(v string) *SettingsUpsert {
	u.Set(settings.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SettingsUpsert) UpdateKey() *SettingsUpsert {
	u.SetExcluded(settings.FieldKey)
	return u
}

// SetValues sets the "values" field.
func (u *SettingsUpsert) SetValues(v []string) *SettingsUpsert {
	u.Set(settings.FieldValues, v)
	return u
}

// UpdateValues sets the "values" field to the value that was provided on create.
func (u *SettingsUpsert) UpdateValues() *SettingsUpsert {
	u.SetExcluded(settings.FieldValues)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingsUpsertOne) UpdateNewValues() *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Settings.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SettingsUpsertOne) Ignore() *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingsUpsertOne) DoNothing() *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingsCreate.OnConflict
// documentation for more info.
func (u *SettingsUpsertOne) Update(set func(*SettingsUpsert)) *SettingsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingsUpsert{UpdateSet: update})
	}))
	return u
}

// SetKey sets the "key" field.
func (u *SettingsUpsertOne) SetKey(v string) *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SettingsUpsertOne) UpdateKey() *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateKey()
	})
}

// SetValues sets the "values" field.
func (u *SettingsUpsertOne) SetValues(v []string) *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.SetValues(v)
	})
}

// UpdateValues sets the "values" field to the value that was provided on create.
func (u *SettingsUpsertOne) UpdateValues() *SettingsUpsertOne {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateValues()
	})
}

// Exec executes the query.
func (u *SettingsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SettingsUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SettingsUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SettingsCreateBulk is the builder for creating many Settings entities in bulk.
type SettingsCreateBulk struct {
	config
	builders []*SettingsCreate
	conflict []sql.ConflictOption
}

// Save creates the Settings entities in the database.
func (scb *SettingsCreateBulk) Save(ctx context.Context) ([]*Settings, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Settings, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingsMutation)
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
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
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
func (scb *SettingsCreateBulk) SaveX(ctx context.Context) []*Settings {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingsCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingsCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Settings.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingsUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
func (scb *SettingsCreateBulk) OnConflict(opts ...sql.ConflictOption) *SettingsUpsertBulk {
	scb.conflict = opts
	return &SettingsUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SettingsCreateBulk) OnConflictColumns(columns ...string) *SettingsUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SettingsUpsertBulk{
		create: scb,
	}
}

// SettingsUpsertBulk is the builder for "upsert"-ing
// a bulk of Settings nodes.
type SettingsUpsertBulk struct {
	create *SettingsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingsUpsertBulk) UpdateNewValues() *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Settings.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SettingsUpsertBulk) Ignore() *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingsUpsertBulk) DoNothing() *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingsCreateBulk.OnConflict
// documentation for more info.
func (u *SettingsUpsertBulk) Update(set func(*SettingsUpsert)) *SettingsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingsUpsert{UpdateSet: update})
	}))
	return u
}

// SetKey sets the "key" field.
func (u *SettingsUpsertBulk) SetKey(v string) *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SettingsUpsertBulk) UpdateKey() *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateKey()
	})
}

// SetValues sets the "values" field.
func (u *SettingsUpsertBulk) SetValues(v []string) *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.SetValues(v)
	})
}

// UpdateValues sets the "values" field to the value that was provided on create.
func (u *SettingsUpsertBulk) UpdateValues() *SettingsUpsertBulk {
	return u.Update(func(s *SettingsUpsert) {
		s.UpdateValues()
	})
}

// Exec executes the query.
func (u *SettingsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SettingsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
