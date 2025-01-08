CREATE TABLE cart_items (
  id VARCHAR(255) PRIMARY KEY,
  customer_id VARCHAR(255) NOT NULL,
  product_id VARCHAR(255) NOT NULL,
  variant_id VARCHAR(255) NOT NULL,
  quantity INTEGER NOT NULL,
  created_at BIGINT,
  updated_at BIGINT,
  UNIQUE(customer_id, product_id, variant_id),
  FOREIGN KEY(customer_id) REFERENCES customers(id),
  FOREIGN KEY(product_id) REFERENCES products(id),
  FOREIGN KEY(variant_id) REFERENCES product_variants(id)
);
