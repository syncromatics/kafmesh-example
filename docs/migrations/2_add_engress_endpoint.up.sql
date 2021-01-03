CREATE TABLE egress_endpoints
    (
        customer_id BIGINT NOT NULL,
        endpoint_url VARCHAR NOT NULL,
        CONSTRAINT UQ_ee_customer_id UNIQUE (customer_id)
    );
