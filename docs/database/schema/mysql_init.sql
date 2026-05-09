CREATE DATABASE IF NOT EXISTS telco_bss;
USE telco_bss;

CREATE TABLE IF NOT EXISTS invoices (
  invoice_id VARCHAR(64) PRIMARY KEY,
  customer_id VARCHAR(64) NOT NULL,
  total_cents BIGINT NOT NULL,
  status VARCHAR(32) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NULL,
  INDEX idx_invoices_customer_id (customer_id)
);

CREATE TABLE IF NOT EXISTS invoice_items (
  invoice_item_id INT AUTO_INCREMENT PRIMARY KEY,
  invoice_id VARCHAR(64) NOT NULL,
  description VARCHAR(255) NOT NULL,
  amount_cents BIGINT NOT NULL,
  created_at DATETIME NOT NULL,
  FOREIGN KEY (invoice_id) REFERENCES invoices(invoice_id)
);

CREATE TABLE IF NOT EXISTS order_items (
  order_item_id INT AUTO_INCREMENT PRIMARY KEY,
  order_id VARCHAR(64) NOT NULL,
  product_code VARCHAR(64) NOT NULL,
  price_cents BIGINT NOT NULL,
  created_at DATETIME NOT NULL,
  INDEX idx_order_items_order_id (order_id)
);

CREATE TABLE IF NOT EXISTS payments (
  payment_id VARCHAR(64) PRIMARY KEY,
  invoice_id VARCHAR(64) NOT NULL,
  amount_cents BIGINT NOT NULL,
  status VARCHAR(32) NOT NULL,
  captured_at DATETIME NOT NULL,
  FOREIGN KEY (invoice_id) REFERENCES invoices(invoice_id)
);
