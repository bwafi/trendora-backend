CREATE TABLE customer_addresses (
  id VARCHAR(255) PRIMARY KEY,
  customer_id VARCHAR(255) NOT NULL,
  recipient_name VARCHAR(100) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  address_type VARCHAR(10) NOT NULL,
  city VARCHAR(100) NOT NULL,
  province VARCHAR(100) NOT NULL,
  sub_district VARCHAR(100) NOT NULL,
  postal_code VARCHAR(10) NOT NULL,
  courier_notes TEXT,
  created_at BIGINT NOT NULL,
  updated_at BIGINT NOT NULL,
  deleted_at TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers (id)
);
