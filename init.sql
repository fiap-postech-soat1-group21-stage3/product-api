CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP DATABASE IF EXISTS restaurant; 
CREATE DATABASE restaurant;
USE restaurant;

-- product
DROP TABLE IF EXISTS product;
CREATE TABLE customer (
	ID  uuid DEFAULT uuid_generate_v4 (),
	NAME VARCHAR(255) NOT NULL,
	DESCRIPTION VARCHAR(255) NOT NULL,
	CATEGORY ENUM('burgers', 'sides', 'beverage', 'sweets') NOT NULL,
	PRICE       DOUBLE PRECISION NOT NULL,
	CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now()
)