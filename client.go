package dent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
	"github.com/jinzhu/copier"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *Schema
	// Dynamic is the client for interacting with the Dynamic builders.
	// Dynamic *DynamicClient
	tmap map[string]*schema.Table
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = NewSchema(c.driver)
	c.tmap = make(map[string]*schema.Table)
	// c.Dynamic = NewDynamicClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		client: c,
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		client: c,
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Dynamic.
//		Query().
//		Count(ctx)
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

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dynamic.Hooks(f(g(h())))`.
func (c *Client) AddTable(table *schema.Table) {
	c.tmap[table.Name] = table
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dynamic.Hooks(f(g(h())))`.
func (c *Client) DeleteTable(name string) {
	delete(c.tmap, name)
}

type Table struct {
	config
	*schema.Table
	client *Client
}

// Table 选择使用哪一个表
func (c *Client) Table(table string) *Table {
	return &Table{
		config: c.config,
		Table:  c.tmap[table],
		client: c,
	}
}

// Create returns a builder for creating a Dynamic entity.
func (c *Table) Create() *DCreate {
	mutation := newDMutation(c.Clone(), OpCreate, withField(c.Columns...))
	return &DCreate{mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dynamic entities.
func (c *Table) CreateBulk(builders ...*DCreate) *DCreateBulk {
	return &DCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dynamic.
func (c *Table) Update() *DUpdate {
	mutation := newDMutation(c.Clone(), OpUpdate)
	return &DUpdate{mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *Table) UpdateOne(d *Dynamic) *DUpdateOne {
	mutation := newDMutation(c.Clone(), OpUpdateOne, withEntity(d))
	return &DUpdateOne{mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *Table) UpdateOneID(id int) *DUpdateOne {
	mutation := newDMutation(c.Clone(), OpUpdateOne, withID(id))
	return &DUpdateOne{mutation: mutation}
}

// Delete returns a delete builder for Dynamic.
func (c *Table) Delete() *DDelete {
	mutation := newDMutation(c.Clone(), OpDelete, withField(c.Columns...))
	return &DDelete{mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *Table) DeleteOneID(id int) *DDeleteOne {
	builder := c.Delete().Where(ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DDeleteOne{builder}
}

// Query returns a query builder for Dynamic.
func (c *Table) Query() *DQuery {
	return &DQuery{
		table:    c,
		withData: make(map[string]*WithQuery),
	}
}

// Get returns a Dynamic entity by its id.
func (c *Table) Get(ctx context.Context, id int) (*Dynamic, error) {
	return c.Query().Where(ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *Table) GetX(ctx context.Context, id int) *Dynamic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *Table) GetColumns() []string {
	var columns []string
	for _, c2 := range c.Columns {
		columns = append(columns, c2.Name)
	}

	return columns
}

// scanValues returns the types for scanning values from sql.Rows.
func (c *Table) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i, v := range columns {
		if val, ok := c.Column(v); ok {
			switch val.Type {
			case field.TypeBool:
				values[i] = new(sql.NullBool)
			case field.TypeTime:
				values[i] = new(sql.NullTime)
			case field.TypeJSON, field.TypeBytes:
				values[i] = new([]byte)
			case field.TypeUUID, field.TypeString:
				values[i] = new(sql.NullString)
			case field.TypeEnum, field.TypeInt8, field.TypeInt16, field.TypeInt32, field.TypeInt, field.TypeInt64, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint, field.TypeUint64:
				values[i] = new(sql.NullInt64)
			case field.TypeFloat32, field.TypeFloat64:
				values[i] = new(sql.NullFloat64)
			case field.TypeOther:
				values[i] = new(interface{})
			default:
				return nil, fmt.Errorf("unexpected column %q for type Dynamic", columns[i])
			}
		}
	}
	return values, nil
}

// Hooks returns the client hooks.
func (c *Table) Clone() *Table {
	var t Table
	copier.Copy(&t, c)
	return &t
}
