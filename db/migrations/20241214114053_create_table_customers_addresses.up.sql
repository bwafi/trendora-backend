create table customer_addresses (
  id VARCHAR(255) PRIMARY KEY,
  user_id VARCHAR(255),
  customer_id VARCHAR(255) NOT NULL,
  recipient_name VARCHAR(255) NOT NULL,
  phone_number varchar(20) NOT NULL,
  address_type varchar[10] not null,
  city varchar(100) NOT NULL,
  province varchar(100) NOT NULL,
  sub_district varchar(100),
  postal_code varchar(10),
  courier_notes text,
  created_at BIGINT NOT NULL,
  updated_at BIGINT NOT NULL,
  deleted_at TIMESTAMP,
  FOREIGN KEY(user_id) REFERENCES customers(id)
);

