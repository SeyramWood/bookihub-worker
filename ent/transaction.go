// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Reference holds the value of the "reference" field.
	Reference string `json:"reference,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// Vat holds the value of the "vat" field.
	Vat float64 `json:"vat,omitempty"`
	// TransactionFee holds the value of the "transaction_fee" field.
	TransactionFee float64 `json:"transaction_fee,omitempty"`
	// CancellationFee holds the value of the "cancellation_fee" field.
	CancellationFee float64 `json:"cancellation_fee,omitempty"`
	// PaidAt holds the value of the "paid_at" field.
	PaidAt time.Time `json:"paid_at,omitempty"`
	// CanceledAt holds the value of the "canceled_at" field.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Channel holds the value of the "channel" field.
	Channel transaction.Channel `json:"channel,omitempty"`
	// TansKind holds the value of the "tans_kind" field.
	TansKind transaction.TansKind `json:"tans_kind,omitempty"`
	// Product holds the value of the "product" field.
	Product transaction.Product `json:"product,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges                TransactionEdges `json:"edges"`
	booking_transaction  *int
	company_transactions *int
	parcel_transaction   *int
	selectValues         sql.SelectValues
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// Booking holds the value of the booking edge.
	Booking *Booking `json:"booking,omitempty"`
	// Parcel holds the value of the parcel edge.
	Parcel *Parcel `json:"parcel,omitempty"`
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BookingOrErr returns the Booking value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) BookingOrErr() (*Booking, error) {
	if e.loadedTypes[0] {
		if e.Booking == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: booking.Label}
		}
		return e.Booking, nil
	}
	return nil, &NotLoadedError{edge: "booking"}
}

// ParcelOrErr returns the Parcel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) ParcelOrErr() (*Parcel, error) {
	if e.loadedTypes[1] {
		if e.Parcel == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: parcel.Label}
		}
		return e.Parcel, nil
	}
	return nil, &NotLoadedError{edge: "parcel"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[2] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldAmount, transaction.FieldVat, transaction.FieldTransactionFee, transaction.FieldCancellationFee:
			values[i] = new(sql.NullFloat64)
		case transaction.FieldID:
			values[i] = new(sql.NullInt64)
		case transaction.FieldReference, transaction.FieldChannel, transaction.FieldTansKind, transaction.FieldProduct:
			values[i] = new(sql.NullString)
		case transaction.FieldCreatedAt, transaction.FieldUpdatedAt, transaction.FieldPaidAt, transaction.FieldCanceledAt:
			values[i] = new(sql.NullTime)
		case transaction.ForeignKeys[0]: // booking_transaction
			values[i] = new(sql.NullInt64)
		case transaction.ForeignKeys[1]: // company_transactions
			values[i] = new(sql.NullInt64)
		case transaction.ForeignKeys[2]: // parcel_transaction
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case transaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case transaction.FieldReference:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference", values[i])
			} else if value.Valid {
				t.Reference = value.String
			}
		case transaction.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				t.Amount = value.Float64
			}
		case transaction.FieldVat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field vat", values[i])
			} else if value.Valid {
				t.Vat = value.Float64
			}
		case transaction.FieldTransactionFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_fee", values[i])
			} else if value.Valid {
				t.TransactionFee = value.Float64
			}
		case transaction.FieldCancellationFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cancellation_fee", values[i])
			} else if value.Valid {
				t.CancellationFee = value.Float64
			}
		case transaction.FieldPaidAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field paid_at", values[i])
			} else if value.Valid {
				t.PaidAt = value.Time
			}
		case transaction.FieldCanceledAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field canceled_at", values[i])
			} else if value.Valid {
				t.CanceledAt = value.Time
			}
		case transaction.FieldChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel", values[i])
			} else if value.Valid {
				t.Channel = transaction.Channel(value.String)
			}
		case transaction.FieldTansKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tans_kind", values[i])
			} else if value.Valid {
				t.TansKind = transaction.TansKind(value.String)
			}
		case transaction.FieldProduct:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product", values[i])
			} else if value.Valid {
				t.Product = transaction.Product(value.String)
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field booking_transaction", value)
			} else if value.Valid {
				t.booking_transaction = new(int)
				*t.booking_transaction = int(value.Int64)
			}
		case transaction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_transactions", value)
			} else if value.Valid {
				t.company_transactions = new(int)
				*t.company_transactions = int(value.Int64)
			}
		case transaction.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field parcel_transaction", value)
			} else if value.Valid {
				t.parcel_transaction = new(int)
				*t.parcel_transaction = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Transaction.
// This includes values selected through modifiers, order, etc.
func (t *Transaction) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryBooking queries the "booking" edge of the Transaction entity.
func (t *Transaction) QueryBooking() *BookingQuery {
	return NewTransactionClient(t.config).QueryBooking(t)
}

// QueryParcel queries the "parcel" edge of the Transaction entity.
func (t *Transaction) QueryParcel() *ParcelQuery {
	return NewTransactionClient(t.config).QueryParcel(t)
}

// QueryCompany queries the "company" edge of the Transaction entity.
func (t *Transaction) QueryCompany() *CompanyQuery {
	return NewTransactionClient(t.config).QueryCompany(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return NewTransactionClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("reference=")
	builder.WriteString(t.Reference)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("vat=")
	builder.WriteString(fmt.Sprintf("%v", t.Vat))
	builder.WriteString(", ")
	builder.WriteString("transaction_fee=")
	builder.WriteString(fmt.Sprintf("%v", t.TransactionFee))
	builder.WriteString(", ")
	builder.WriteString("cancellation_fee=")
	builder.WriteString(fmt.Sprintf("%v", t.CancellationFee))
	builder.WriteString(", ")
	builder.WriteString("paid_at=")
	builder.WriteString(t.PaidAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("canceled_at=")
	builder.WriteString(t.CanceledAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("channel=")
	builder.WriteString(fmt.Sprintf("%v", t.Channel))
	builder.WriteString(", ")
	builder.WriteString("tans_kind=")
	builder.WriteString(fmt.Sprintf("%v", t.TansKind))
	builder.WriteString(", ")
	builder.WriteString("product=")
	builder.WriteString(fmt.Sprintf("%v", t.Product))
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction
