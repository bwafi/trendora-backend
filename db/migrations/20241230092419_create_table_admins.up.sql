CREATE TABLE admins (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    refresh_token VARCHAR(255),
    created_at bigint,
    updated_at bigint,
    deleted_at bigint
);
