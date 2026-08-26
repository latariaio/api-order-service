ALTER TABLE customers ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE customers ADD COLUMN deleted_at TIMESTAMP;

CREATE INDEX idx_customers_deleted_at ON customers (deleted_at);