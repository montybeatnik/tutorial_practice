CREATE TABLE IF NOT EXISTS hardware (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    vendor text,
    model text,
    UNIQUE (vendor, model)
);