-- +goose Up
CREATE TABLE Feed (
    user_feed UUID REFERENCES Registeredusers(id),
    disable BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE Feed;