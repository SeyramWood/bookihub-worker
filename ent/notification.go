// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/notification"
	"github.com/SeyramWood/ent/schema"
)

// Notification is the model entity for the Notification schema.
type Notification struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Event holds the value of the "event" field.
	Event string `json:"event,omitempty"`
	// Activity holds the value of the "activity" field.
	Activity string `json:"activity,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// SubjectType holds the value of the "subject_type" field.
	SubjectType string `json:"subject_type,omitempty"`
	// SubjectID holds the value of the "subject_id" field.
	SubjectID int `json:"subject_id,omitempty"`
	// CreatorType holds the value of the "creator_type" field.
	CreatorType string `json:"creator_type,omitempty"`
	// CreatorID holds the value of the "creator_id" field.
	CreatorID int `json:"creator_id,omitempty"`
	// CustomerReadAt holds the value of the "customer_read_at" field.
	CustomerReadAt string `json:"customer_read_at,omitempty"`
	// BookibusReadAt holds the value of the "bookibus_read_at" field.
	BookibusReadAt []*schema.NotificationRead `json:"bookibus_read_at,omitempty"`
	// CompanyReadAt holds the value of the "company_read_at" field.
	CompanyReadAt []*schema.NotificationRead `json:"company_read_at,omitempty"`
	// Data holds the value of the "data" field.
	Data *struct {
		Data interface{} "json:\"data\""
	} `json:"data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationQuery when eager-loading is set.
	Edges                 NotificationEdges `json:"edges"`
	company_notifications *int
	selectValues          sql.SelectValues
}

// NotificationEdges holds the relations/edges for other nodes in the graph.
type NotificationEdges struct {
	// BookibusUser holds the value of the bookibus_user edge.
	BookibusUser []*BookibusUser `json:"bookibus_user,omitempty"`
	// CompanyUser holds the value of the company_user edge.
	CompanyUser []*CompanyUser `json:"company_user,omitempty"`
	// Customer holds the value of the customer edge.
	Customer []*Customer `json:"customer,omitempty"`
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// BookibusUserOrErr returns the BookibusUser value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) BookibusUserOrErr() ([]*BookibusUser, error) {
	if e.loadedTypes[0] {
		return e.BookibusUser, nil
	}
	return nil, &NotLoadedError{edge: "bookibus_user"}
}

// CompanyUserOrErr returns the CompanyUser value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) CompanyUserOrErr() ([]*CompanyUser, error) {
	if e.loadedTypes[1] {
		return e.CompanyUser, nil
	}
	return nil, &NotLoadedError{edge: "company_user"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) CustomerOrErr() ([]*Customer, error) {
	if e.loadedTypes[2] {
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotificationEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[3] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notification) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notification.FieldBookibusReadAt, notification.FieldCompanyReadAt, notification.FieldData:
			values[i] = new([]byte)
		case notification.FieldID, notification.FieldSubjectID, notification.FieldCreatorID:
			values[i] = new(sql.NullInt64)
		case notification.FieldEvent, notification.FieldActivity, notification.FieldDescription, notification.FieldSubjectType, notification.FieldCreatorType, notification.FieldCustomerReadAt:
			values[i] = new(sql.NullString)
		case notification.ForeignKeys[0]: // company_notifications
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notification fields.
func (n *Notification) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case notification.FieldEvent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event", values[i])
			} else if value.Valid {
				n.Event = value.String
			}
		case notification.FieldActivity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field activity", values[i])
			} else if value.Valid {
				n.Activity = value.String
			}
		case notification.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				n.Description = value.String
			}
		case notification.FieldSubjectType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject_type", values[i])
			} else if value.Valid {
				n.SubjectType = value.String
			}
		case notification.FieldSubjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field subject_id", values[i])
			} else if value.Valid {
				n.SubjectID = int(value.Int64)
			}
		case notification.FieldCreatorType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_type", values[i])
			} else if value.Valid {
				n.CreatorType = value.String
			}
		case notification.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				n.CreatorID = int(value.Int64)
			}
		case notification.FieldCustomerReadAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_read_at", values[i])
			} else if value.Valid {
				n.CustomerReadAt = value.String
			}
		case notification.FieldBookibusReadAt:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field bookibus_read_at", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.BookibusReadAt); err != nil {
					return fmt.Errorf("unmarshal field bookibus_read_at: %w", err)
				}
			}
		case notification.FieldCompanyReadAt:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field company_read_at", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.CompanyReadAt); err != nil {
					return fmt.Errorf("unmarshal field company_read_at: %w", err)
				}
			}
		case notification.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		case notification.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_notifications", value)
			} else if value.Valid {
				n.company_notifications = new(int)
				*n.company_notifications = int(value.Int64)
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Notification.
// This includes values selected through modifiers, order, etc.
func (n *Notification) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryBookibusUser queries the "bookibus_user" edge of the Notification entity.
func (n *Notification) QueryBookibusUser() *BookibusUserQuery {
	return NewNotificationClient(n.config).QueryBookibusUser(n)
}

// QueryCompanyUser queries the "company_user" edge of the Notification entity.
func (n *Notification) QueryCompanyUser() *CompanyUserQuery {
	return NewNotificationClient(n.config).QueryCompanyUser(n)
}

// QueryCustomer queries the "customer" edge of the Notification entity.
func (n *Notification) QueryCustomer() *CustomerQuery {
	return NewNotificationClient(n.config).QueryCustomer(n)
}

// QueryCompany queries the "company" edge of the Notification entity.
func (n *Notification) QueryCompany() *CompanyQuery {
	return NewNotificationClient(n.config).QueryCompany(n)
}

// Update returns a builder for updating this Notification.
// Note that you need to call Notification.Unwrap() before calling this method if this Notification
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notification) Update() *NotificationUpdateOne {
	return NewNotificationClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Notification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notification) Unwrap() *Notification {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Notification is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notification) String() string {
	var builder strings.Builder
	builder.WriteString("Notification(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("event=")
	builder.WriteString(n.Event)
	builder.WriteString(", ")
	builder.WriteString("activity=")
	builder.WriteString(n.Activity)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(n.Description)
	builder.WriteString(", ")
	builder.WriteString("subject_type=")
	builder.WriteString(n.SubjectType)
	builder.WriteString(", ")
	builder.WriteString("subject_id=")
	builder.WriteString(fmt.Sprintf("%v", n.SubjectID))
	builder.WriteString(", ")
	builder.WriteString("creator_type=")
	builder.WriteString(n.CreatorType)
	builder.WriteString(", ")
	builder.WriteString("creator_id=")
	builder.WriteString(fmt.Sprintf("%v", n.CreatorID))
	builder.WriteString(", ")
	builder.WriteString("customer_read_at=")
	builder.WriteString(n.CustomerReadAt)
	builder.WriteString(", ")
	builder.WriteString("bookibus_read_at=")
	builder.WriteString(fmt.Sprintf("%v", n.BookibusReadAt))
	builder.WriteString(", ")
	builder.WriteString("company_read_at=")
	builder.WriteString(fmt.Sprintf("%v", n.CompanyReadAt))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", n.Data))
	builder.WriteByte(')')
	return builder.String()
}

// Notifications is a parsable slice of Notification.
type Notifications []*Notification