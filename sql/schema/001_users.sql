-- +goose Up
CREATE TABLE IF NOT EXISTS Unregisteredusers(
    whatsapp_number VARCHAR(20) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    display_name TEXT NOT NULL
); 

CREATE TABLE IF NOT EXISTS Registeredusers(
    id UUID PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email VARCHAR(64) NOT NULL ,
    isEmailVerified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP NOT NULL,
    password VARCHAR(64) NOT NULL,
    whatsapp_number VARCHAR(20) NOT NULL,
    loggedIn BOOLEAN DEFAULT FALSE,
    apikey VARCHAR(64) NOT NULL,
    updated_username BOOLEAN DEFAULT FALSE,
    updated_password BOOLEAN DEFAULT FALSE,
    UNIQUE(email, whatsapp_number)
);
-- +goose Down
DROP TABLE Unregisteredusers;
DROP TABLE Registeredusers;