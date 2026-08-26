DROP INDEX IF EXISTS idx_customers_document;
DROP INDEX IF EXISTS idx_customers_email;

CREATE UNIQUE INDEX idx_customers_document ON customers (document) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_customers_email ON customers (email) WHERE deleted_at IS NULL;