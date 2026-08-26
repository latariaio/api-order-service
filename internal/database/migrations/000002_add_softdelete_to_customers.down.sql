DROP INDEX IF EXISTS idx_customers_deleted_at;
ALTER TABLE customers DROP COLUMN updated_at;
ALTER TABLE customers DROP COLUMN deleted_at;