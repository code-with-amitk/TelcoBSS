# Couchbase Setup

This file documents Couchbase setup for TelcoBSS.

## Bucket and Collections

Create a bucket named `telco_bss` with default scope and the following collections:

- `customers`
- `orders`
- `usage_records`

## Indexes

Create a primary index on the bucket:

```
CREATE PRIMARY INDEX ON `telco_bss`;
```

Create secondary indexes:

```
CREATE INDEX idx_customers_customer_id ON `telco_bss`.`_default`.`customers`(customer_id);
CREATE INDEX idx_orders_order_date ON `telco_bss`.`_default`.`orders`(order_date);
```

## Go Interaction

In Go, use `gocb/v2`:

- `Connect` to the cluster.
- `Bucket("telco_bss")` and `DefaultScope().Collection("customers")`.
- Use `Upsert`, `Get`, and `Query` with N1QL.

Example repository methods are in `internal/common/db/couchbase.go`.

## C++ Interaction

In C++, use `libcouchbase` or `couchbase++`.

Example placeholders are available in `cmd/rating-engine/main.cpp` with comments showing how Kafka and Couchbase would be wired together.
