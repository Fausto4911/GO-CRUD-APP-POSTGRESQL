CREATE TABLE IF NOT EXISTS client (
    id SERIAL PRIMARY KEY NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	nick_name VARCHAR(255),
	create_at TIMESTAMP DEFAULT NOW()
	update_at TIMESTAMP
);
