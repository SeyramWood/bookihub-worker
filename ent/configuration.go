// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/bookibus/ent/configuration"
	"github.com/SeyramWood/bookibus/ent/schema"
)

// Configuration is the model entity for the Configuration schema.
type Configuration struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Charge holds the value of the "charge" field.
	Charge       *schema.Charge `json:"charge,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Configuration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case configuration.FieldCharge:
			values[i] = new([]byte)
		case configuration.FieldID:
			values[i] = new(sql.NullInt64)
		case configuration.FieldCreatedAt, configuration.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Configuration fields.
func (c *Configuration) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case configuration.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case configuration.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case configuration.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case configuration.FieldCharge:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field charge", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Charge); err != nil {
					return fmt.Errorf("unmarshal field charge: %w", err)
				}
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Configuration.
// This includes values selected through modifiers, order, etc.
func (c *Configuration) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Configuration.
// Note that you need to call Configuration.Unwrap() before calling this method if this Configuration
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Configuration) Update() *ConfigurationUpdateOne {
	return NewConfigurationClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Configuration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Configuration) Unwrap() *Configuration {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Configuration is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Configuration) String() string {
	var builder strings.Builder
	builder.WriteString("Configuration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("charge=")
	builder.WriteString(fmt.Sprintf("%v", c.Charge))
	builder.WriteByte(')')
	return builder.String()
}

// Configurations is a parsable slice of Configuration.
type Configurations []*Configuration
