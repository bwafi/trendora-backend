CREATE TABLE customers (
  id VARCHAR(255) PRIMARY KEY,
  email_address VARCHAR(255) UNIQUE NOT NULL,
  phone_number VARCHAR(20) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  gender BOOLEAN NOT NULL,
  date_of_birth BIGINT NOT NULL,
  password VARCHAR(255) NOT NULL,
  token VARCHAR(255),
  created_at BIGINT NOT NULL,
  updated_at BIGINT NOT NULL,
  deleted_at TIMESTAMP
);
