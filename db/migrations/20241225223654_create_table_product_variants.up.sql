CREATE TABLE product_variants (
  id VARCHAR(255) PRIMARY KEY,
  product_id VARCHAR(255),
  sku VARCHAR(100),
  color_name VARCHAR(50),
  size VARCHAR(20),
  discount DECIMAL(5,2),
  price DECIMAL(10,2),
  stock_quantity INTEGER,
  weight DECIMAL(8,2),
  is_available BOOLEAN,
  FOREIGN KEY(product_id) REFERENCES products(id)
)
