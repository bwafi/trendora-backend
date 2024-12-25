CREATE TABLE categories (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  slug VARCHAR(100) NOT NULL,
  parent_id VARCHAR(255),
  FOREIGN KEY(parent_id) REFERENCES categories(id)
);
