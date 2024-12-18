CREATE TABLE customer_sessions (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  customer_id VARCHAR(255) NOT NULL,
  refresh_token VARCHAR(512) NOT NULL,
  is_revoked BOOL NOT NULL DEFAULT false,
  created_at BIGINT NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers (id)
);
