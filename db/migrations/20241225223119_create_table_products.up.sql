CREATE TABLE products (
  id VARCHAR(255) PRIMARY KEY,
  style_code VARCHAR(50) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  category_id VARCHAR(255),
  sub_category_id VARCHAR(255),
  base_price DECIMAL(10,2),
  is_visible BOOLEAN,
  release_date BIGINT,
  created_at BIGINT,
  updated_at BIGINT,
  FOREIGN KEY(category_id) REFERENCES categories(id),
  FOREIGN KEY(sub_category_id) REFERENCES categories(id)
)
