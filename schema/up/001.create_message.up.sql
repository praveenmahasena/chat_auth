CREATE TABLE IF NOT EXISTS messages(
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id)
);
