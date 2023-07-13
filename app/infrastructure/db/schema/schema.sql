CREATE TABLE users (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  phone_number VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  postal_code VARCHAR(255) NOT NULL COMMENT "郵便番号",
  prefecture VARCHAR(255) NOT NULL COMMENT "都道府県",
  city VARCHAR(255) NOT NULL COMMENT "市区町村",
  address VARCHAR(255) NOT NULL COMMENT "番地",
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);