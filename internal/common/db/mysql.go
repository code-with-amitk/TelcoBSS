package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping mysql: %w", err)
	}

	return db, nil
}

// InsertInvoice writes an invoice record into MySQL.
func InsertInvoice(db *sql.DB, invoiceID string, customerID string, totalCents int64) error {
	query := `INSERT INTO invoices (invoice_id, customer_id, total_cents, status, created_at) VALUES (?, ?, ?, ?, NOW())`
	_, err := db.Exec(query, invoiceID, customerID, totalCents, "OPEN")
	return err
}

// InsertOrderItem writes an order item into MySQL.
func InsertOrderItem(db *sql.DB, orderID string, productCode string, priceCents int64) error {
	query := `INSERT INTO order_items (order_id, product_code, price_cents, created_at) VALUES (?, ?, ?, NOW())`
	_, err := db.Exec(query, orderID, productCode, priceCents)
	return err
}
