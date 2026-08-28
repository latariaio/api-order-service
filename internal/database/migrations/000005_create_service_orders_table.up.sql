CREATE SEQUENCE service_order_number_seq START 100;

CREATE TABLE service_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_number BIGINT NOT NULL DEFAULT nextval('service_order_number_seq'),
    customer_id UUID NOT NULL REFERENCES customers(id),
    reported_problem TEXT,
    status VARCHAR(30) NOT NULL DEFAULT 'ABERTA',
    total_price NUMERIC(10,2) NOT NULL DEFAULT 0,
    diagnosis TEXT,
    notes TEXT,
    opened_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_service_orders_order_number ON service_orders (order_number);
CREATE INDEX idx_service_orders_customer_id ON service_orders (customer_id);
CREATE INDEX idx_service_orders_status ON service_orders (status);