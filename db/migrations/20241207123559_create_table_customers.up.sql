create table customers (
  id VARCHAR(255) PRIMARY KEY,
  email_address VARCHAR(255) UNIQUE,
  phone_number VARCHAR(20) UNIQUE,
  name VARCHAR(50) NOT NULL,
  gender BOOL,
  date_of_birth BIGINT,
  token VARCHAR(255),
  password VARCHAR(255) NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT NOT NULL,
  deleted_at TIMESTAMP
);
