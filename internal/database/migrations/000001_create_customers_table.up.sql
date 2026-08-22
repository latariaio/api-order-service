CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    document VARCHAR(50) NOT NULL,
    document_type VARCHAR(20) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255),
    address_street VARCHAR(255) NOT NULL,
    address_number VARCHAR(20) NOT NULL,
    address_complement VARCHAR(255),
    address_neighborhood VARCHAR(255) NOT NULL,
    address_city VARCHAR(255) NOT NULL,
    address_state VARCHAR(2) NOT NULL,
    address_zip_code VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX idx_customers_document ON customers (document);
CREATE UNIQUE INDEX idx_customers_email ON customers (email);