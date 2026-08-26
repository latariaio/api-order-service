DROP INDEX IF EXISTS idx_customers_document;
DROP INDEX IF EXISTS idx_customers_email;

CREATE UNIQUE INDEX idx_customers_document ON customers (document);
CREATE UNIQUE INDEX idx_customers_email ON customers (email);