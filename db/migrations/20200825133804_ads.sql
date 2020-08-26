
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Ads
(
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(200) NOT NULL,
    Description VARCHAR(1000) NOT NULL,
    IMG1 VARCHAR NOT NULL,
    IMG2 VARCHAR,
    IMG3 VARCHAR,
    Price NUMERIC NOT NULL,

    Created_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Ads;
