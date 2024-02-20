-- +goose Up
CREATE TYPE subscription_status_enum AS ENUM ('Active', 'Expired');
CREATE TYPE subscription_tier_enum AS ENUM('Basic', 'Pro', 'Free-Trial');
CREATE TYPE payment_status_enum AS ENUM('Successful', 'Pending', 'Failed');
CREATE TABLE Subscription(
    subscription_id UUID PRIMARY KEY,
    userid UUID UNIQUE,
    expiry_date DATE NOT NULL,
    whatsapp_number VARCHAR(20) NOT NULL,
    status subscription_status_enum DEFAULT 'Active',
    registered_user BOOLEAN DEFAULT FALSE,
    subscription_tier subscription_tier_enum DEFAULT 'Free-Trial',
    FOREIGN KEY (userid) REFERENCES RegisteredUsers(id)
);

CREATE TABLE Plan (
    plan_id SERIAL PRIMARY KEY,
    plan_name VARCHAR(16) NOT NULL UNIQUE,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL UNIQUE,
    duration_days INT NOT NULL
);

CREATE TABLE Payment (
    payment_id UUID PRIMARY KEY,
    userid UUID,
    amount DECIMAL(10, 2) REFERENCES Plan(price),
    plan_id INT REFERENCES Plan(plan_id),
    payment_date DATE,
    expiry_date DATE NOT NULL,
    payment_method VARCHAR(50),
    payment_status payment_status_enum DEFAULT 'Pending',
    FOREIGN KEY (userid) REFERENCES Subscriptions(userid)
);

-- +goose Down
DROP TABLE Subscriptions;
DROP TABLE Plan;
DROP TABLE Payment;
DROP TYPE payment_status_enum;
DROP TYPE subscription_status_enum;