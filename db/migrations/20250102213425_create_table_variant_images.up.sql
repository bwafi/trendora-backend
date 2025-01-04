CREATE TABLE variant_images (
  id VARCHAR(255) PRIMARY KEY,
  variant_id VARCHAR(255),
  image_url VARCHAR(500) NOT NULL,
  display_order INTEGER,
  FOREIGN KEY(variant_id) REFERENCES product_variants(id)
)
