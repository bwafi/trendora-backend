CREATE TABLE product_variants (
  id VARCHAR(255) PRIMARY KEY,
  product_id VARCHAR(255),
  sku VARCHAR(100) NOT NULL UNIQUE,
  color_name VARCHAR(50),
  weight DECIMAL(8,2),
  is_available BOOLEAN,
  FOREIGN KEY(product_id) REFERENCES products(id)
)
