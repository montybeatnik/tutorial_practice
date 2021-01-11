CREATE TABLE IF NOT EXISTS devices (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    hostname text UNIQUE,
    loopback text,
    hardware_id INTEGER REFERENCES hardware (id),
    software_id INTEGER REFERENCES software (id)
);