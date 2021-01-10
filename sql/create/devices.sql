CREATE TABLE IF NOT EXISTS devices (
    id SERIAL PRIMARY KEY,
    hostname text,
    loopback text,
    models_id INTEGER REFERENCES models (id),
    software_version_id INTEGER REFERENCES software_versions (id)
);