-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Asia/Jakarta";

-- Create users table
CREATE TABLE users (
                     id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
                     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                     updated_at TIMESTAMP NULL,
                     email VARCHAR (50) NOT NULL UNIQUE,
                     username VARCHAR (50) NOT NULL UNIQUE,
                     password_hash VARCHAR (255) NOT NULL,
                     user_status INT NOT NULL,
                     user_role VARCHAR (30) NOT NULL,
                     display_name VARCHAR (50) NOT NULL
);

-- Add indexes
CREATE INDEX active_users ON users (id) WHERE user_status = 1;
CREATE INDEX active_users ON users (username) WHERE user_status = 1;
CREATE INDEX active_users ON users (display_name) WHERE user_status = 1;