CREATE TABLE customers
(
  customer_id BIGINT,
  name VARCHAR NOT NULL
);

INSERT INTO customers (customer_id, name) VALUES
  (1, "test customer"),
  (2, "second customer");
