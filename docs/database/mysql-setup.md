# MySQL Setup

TelcoBSS uses MySQL for relational billing and order data.

## Schema

The primary tables are:

- `invoices`
- `invoice_items`
- `order_items`
- `payments`

Foreign keys ensure data integrity between orders, invoices, and payments.

## Example Table Relationships

- `order_items` references `invoices` via `order_id`.
- `invoice_items` references `invoices` via `invoice_id`.

## Go Interaction

In Go, use `database/sql` with the `go-sql-driver/mysql` driver.

Example methods are in `internal/common/db/mysql.go`.

## Initialization

Run the SQL schema from `docs/database/schema/mysql_init.sql` to create tables and indexes.
