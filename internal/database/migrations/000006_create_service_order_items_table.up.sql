CREATE TABLE service_order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_order_id UUID NOT NULL REFERENCES service_orders(id),
    service_id UUID NOT NULL REFERENCES services(id),
    quantity INTEGER NOT NULL,
    unit_price NUMERIC(10,2) NOT NULL,
    total_price NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_service_order_items_service_order_id ON service_order_items (service_order_id);
CREATE INDEX idx_service_order_items_service_id ON service_order_items (service_id);
CREATE INDEX idx_service_order_items_deleted_at ON service_order_items (deleted_at);