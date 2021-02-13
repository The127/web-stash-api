// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"web-stash-api/ent/migrate"

	"web-stash-api/ent/bag"
	"web-stash-api/ent/bagitem"
	"web-stash-api/ent/share"
	"web-stash-api/ent/subitem"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Bag is the client for interacting with the Bag builders.
	Bag *BagClient
	// BagItem is the client for interacting with the BagItem builders.
	BagItem *BagItemClient
	// Share is the client for interacting with the Share builders.
	Share *ShareClient
	// SubItem is the client for interacting with the SubItem builders.
	SubItem *SubItemClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Bag = NewBagClient(c.config)
	c.BagItem = NewBagItemClient(c.config)
	c.Share = NewShareClient(c.config)
	c.SubItem = NewSubItemClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Bag:     NewBagClient(cfg),
		BagItem: NewBagItemClient(cfg),
		Share:   NewShareClient(cfg),
		SubItem: NewSubItemClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:  cfg,
		Bag:     NewBagClient(cfg),
		BagItem: NewBagItemClient(cfg),
		Share:   NewShareClient(cfg),
		SubItem: NewSubItemClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Bag.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Bag.Use(hooks...)
	c.BagItem.Use(hooks...)
	c.Share.Use(hooks...)
	c.SubItem.Use(hooks...)
}

// BagClient is a client for the Bag schema.
type BagClient struct {
	config
}

// NewBagClient returns a client for the Bag from the given config.
func NewBagClient(c config) *BagClient {
	return &BagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bag.Hooks(f(g(h())))`.
func (c *BagClient) Use(hooks ...Hook) {
	c.hooks.Bag = append(c.hooks.Bag, hooks...)
}

// Create returns a create builder for Bag.
func (c *BagClient) Create() *BagCreate {
	mutation := newBagMutation(c.config, OpCreate)
	return &BagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Bag entities.
func (c *BagClient) CreateBulk(builders ...*BagCreate) *BagCreateBulk {
	return &BagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bag.
func (c *BagClient) Update() *BagUpdate {
	mutation := newBagMutation(c.config, OpUpdate)
	return &BagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BagClient) UpdateOne(b *Bag) *BagUpdateOne {
	mutation := newBagMutation(c.config, OpUpdateOne, withBag(b))
	return &BagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BagClient) UpdateOneID(id uuid.UUID) *BagUpdateOne {
	mutation := newBagMutation(c.config, OpUpdateOne, withBagID(id))
	return &BagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bag.
func (c *BagClient) Delete() *BagDelete {
	mutation := newBagMutation(c.config, OpDelete)
	return &BagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BagClient) DeleteOne(b *Bag) *BagDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BagClient) DeleteOneID(id uuid.UUID) *BagDeleteOne {
	builder := c.Delete().Where(bag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BagDeleteOne{builder}
}

// Query returns a query builder for Bag.
func (c *BagClient) Query() *BagQuery {
	return &BagQuery{config: c.config}
}

// Get returns a Bag entity by its id.
func (c *BagClient) Get(ctx context.Context, id uuid.UUID) (*Bag, error) {
	return c.Query().Where(bag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BagClient) GetX(ctx context.Context, id uuid.UUID) *Bag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryItems queries the items edge of a Bag.
func (c *BagClient) QueryItems(b *Bag) *BagItemQuery {
	query := &BagItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bag.Table, bag.FieldID, id),
			sqlgraph.To(bagitem.Table, bagitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bag.ItemsTable, bag.ItemsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BagClient) Hooks() []Hook {
	return c.hooks.Bag
}

// BagItemClient is a client for the BagItem schema.
type BagItemClient struct {
	config
}

// NewBagItemClient returns a client for the BagItem from the given config.
func NewBagItemClient(c config) *BagItemClient {
	return &BagItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bagitem.Hooks(f(g(h())))`.
func (c *BagItemClient) Use(hooks ...Hook) {
	c.hooks.BagItem = append(c.hooks.BagItem, hooks...)
}

// Create returns a create builder for BagItem.
func (c *BagItemClient) Create() *BagItemCreate {
	mutation := newBagItemMutation(c.config, OpCreate)
	return &BagItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BagItem entities.
func (c *BagItemClient) CreateBulk(builders ...*BagItemCreate) *BagItemCreateBulk {
	return &BagItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BagItem.
func (c *BagItemClient) Update() *BagItemUpdate {
	mutation := newBagItemMutation(c.config, OpUpdate)
	return &BagItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BagItemClient) UpdateOne(bi *BagItem) *BagItemUpdateOne {
	mutation := newBagItemMutation(c.config, OpUpdateOne, withBagItem(bi))
	return &BagItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BagItemClient) UpdateOneID(id uuid.UUID) *BagItemUpdateOne {
	mutation := newBagItemMutation(c.config, OpUpdateOne, withBagItemID(id))
	return &BagItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BagItem.
func (c *BagItemClient) Delete() *BagItemDelete {
	mutation := newBagItemMutation(c.config, OpDelete)
	return &BagItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BagItemClient) DeleteOne(bi *BagItem) *BagItemDeleteOne {
	return c.DeleteOneID(bi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BagItemClient) DeleteOneID(id uuid.UUID) *BagItemDeleteOne {
	builder := c.Delete().Where(bagitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BagItemDeleteOne{builder}
}

// Query returns a query builder for BagItem.
func (c *BagItemClient) Query() *BagItemQuery {
	return &BagItemQuery{config: c.config}
}

// Get returns a BagItem entity by its id.
func (c *BagItemClient) Get(ctx context.Context, id uuid.UUID) (*BagItem, error) {
	return c.Query().Where(bagitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BagItemClient) GetX(ctx context.Context, id uuid.UUID) *BagItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBag queries the bag edge of a BagItem.
func (c *BagItemClient) QueryBag(bi *BagItem) *BagQuery {
	query := &BagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bagitem.Table, bagitem.FieldID, id),
			sqlgraph.To(bag.Table, bag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bagitem.BagTable, bagitem.BagColumn),
		)
		fromV = sqlgraph.Neighbors(bi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubItems queries the sub_items edge of a BagItem.
func (c *BagItemClient) QuerySubItems(bi *BagItem) *SubItemQuery {
	query := &SubItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bagitem.Table, bagitem.FieldID, id),
			sqlgraph.To(subitem.Table, subitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bagitem.SubItemsTable, bagitem.SubItemsColumn),
		)
		fromV = sqlgraph.Neighbors(bi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BagItemClient) Hooks() []Hook {
	return c.hooks.BagItem
}

// ShareClient is a client for the Share schema.
type ShareClient struct {
	config
}

// NewShareClient returns a client for the Share from the given config.
func NewShareClient(c config) *ShareClient {
	return &ShareClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `share.Hooks(f(g(h())))`.
func (c *ShareClient) Use(hooks ...Hook) {
	c.hooks.Share = append(c.hooks.Share, hooks...)
}

// Create returns a create builder for Share.
func (c *ShareClient) Create() *ShareCreate {
	mutation := newShareMutation(c.config, OpCreate)
	return &ShareCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Share entities.
func (c *ShareClient) CreateBulk(builders ...*ShareCreate) *ShareCreateBulk {
	return &ShareCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Share.
func (c *ShareClient) Update() *ShareUpdate {
	mutation := newShareMutation(c.config, OpUpdate)
	return &ShareUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ShareClient) UpdateOne(s *Share) *ShareUpdateOne {
	mutation := newShareMutation(c.config, OpUpdateOne, withShare(s))
	return &ShareUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ShareClient) UpdateOneID(id uuid.UUID) *ShareUpdateOne {
	mutation := newShareMutation(c.config, OpUpdateOne, withShareID(id))
	return &ShareUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Share.
func (c *ShareClient) Delete() *ShareDelete {
	mutation := newShareMutation(c.config, OpDelete)
	return &ShareDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ShareClient) DeleteOne(s *Share) *ShareDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ShareClient) DeleteOneID(id uuid.UUID) *ShareDeleteOne {
	builder := c.Delete().Where(share.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ShareDeleteOne{builder}
}

// Query returns a query builder for Share.
func (c *ShareClient) Query() *ShareQuery {
	return &ShareQuery{config: c.config}
}

// Get returns a Share entity by its id.
func (c *ShareClient) Get(ctx context.Context, id uuid.UUID) (*Share, error) {
	return c.Query().Where(share.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ShareClient) GetX(ctx context.Context, id uuid.UUID) *Share {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ShareClient) Hooks() []Hook {
	return c.hooks.Share
}

// SubItemClient is a client for the SubItem schema.
type SubItemClient struct {
	config
}

// NewSubItemClient returns a client for the SubItem from the given config.
func NewSubItemClient(c config) *SubItemClient {
	return &SubItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subitem.Hooks(f(g(h())))`.
func (c *SubItemClient) Use(hooks ...Hook) {
	c.hooks.SubItem = append(c.hooks.SubItem, hooks...)
}

// Create returns a create builder for SubItem.
func (c *SubItemClient) Create() *SubItemCreate {
	mutation := newSubItemMutation(c.config, OpCreate)
	return &SubItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SubItem entities.
func (c *SubItemClient) CreateBulk(builders ...*SubItemCreate) *SubItemCreateBulk {
	return &SubItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SubItem.
func (c *SubItemClient) Update() *SubItemUpdate {
	mutation := newSubItemMutation(c.config, OpUpdate)
	return &SubItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubItemClient) UpdateOne(si *SubItem) *SubItemUpdateOne {
	mutation := newSubItemMutation(c.config, OpUpdateOne, withSubItem(si))
	return &SubItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubItemClient) UpdateOneID(id uuid.UUID) *SubItemUpdateOne {
	mutation := newSubItemMutation(c.config, OpUpdateOne, withSubItemID(id))
	return &SubItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SubItem.
func (c *SubItemClient) Delete() *SubItemDelete {
	mutation := newSubItemMutation(c.config, OpDelete)
	return &SubItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubItemClient) DeleteOne(si *SubItem) *SubItemDeleteOne {
	return c.DeleteOneID(si.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubItemClient) DeleteOneID(id uuid.UUID) *SubItemDeleteOne {
	builder := c.Delete().Where(subitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubItemDeleteOne{builder}
}

// Query returns a query builder for SubItem.
func (c *SubItemClient) Query() *SubItemQuery {
	return &SubItemQuery{config: c.config}
}

// Get returns a SubItem entity by its id.
func (c *SubItemClient) Get(ctx context.Context, id uuid.UUID) (*SubItem, error) {
	return c.Query().Where(subitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubItemClient) GetX(ctx context.Context, id uuid.UUID) *SubItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a SubItem.
func (c *SubItemClient) QueryParent(si *SubItem) *BagItemQuery {
	query := &BagItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := si.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subitem.Table, subitem.FieldID, id),
			sqlgraph.To(bagitem.Table, bagitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subitem.ParentTable, subitem.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(si.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubItemClient) Hooks() []Hook {
	return c.hooks.SubItem
}
