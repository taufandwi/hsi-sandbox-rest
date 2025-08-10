-- Create the 'users' table to store login information.
CREATE TABLE users (
    -- The primary key for the users table.
    -- It uses SERIAL for an auto-incrementing integer ID.
    id SERIAL PRIMARY KEY,

    -- The username for the user. It must be unique across all users
    -- and cannot be empty.
    username VARCHAR(255) UNIQUE NOT NULL,

    -- The hashed password for security.
    -- NEVER store plain text passwords.
    -- The length is set to accommodate a strong hash like bcrypt.
    password_hash VARCHAR(255) NOT NULL,

    -- A timestamp for when the user account was created.
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'employees' table to store detailed employee information.
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE NOT NULL REFERENCES public.users(id),
    -- Personal and professional details of the employee.
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(20),
    hire_date DATE NOT NULL,
    job_title VARCHAR(255),
    department VARCHAR(255)
);