CREATE TABLE software (
    id SERIAL PRIMARY KEY,
    version text UNIQUE,
    created_at timestamp
);