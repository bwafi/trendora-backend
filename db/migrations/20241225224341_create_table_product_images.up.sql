CREATE TABLE product_images (
  id VARCHAR(255) PRIMARY KEY,
  product_id VARCHAR(255),
  variant_id VARCHAR(255),
  image_url VARCHAR(500) NOT NULL,
  display_order INTEGER,
  FOREIGN KEY(product_id) REFERENCES products(id),
  FOREIGN KEY(variant_id) REFERENCES product_variants(id)
)
