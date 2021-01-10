CREATE TABLE software_versions (
    id SERIAL PRIMARY KEY,
    version text UNIQUE,
    created_at timestamp
);