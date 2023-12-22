CREATE TABLE IF NOT EXISTS secret_santa.secret_santas (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL UNIQUE,
    year INT NOT NULL,
    group_id INT NOT NULL,
    secret_santa_id INT NOT NULL,
    recipient_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (secret_santa_id) REFERENCES secret_santa.users(id) ON DELETE CASCADE,
    FOREIGN KEY (recipient_id) REFERENCES secret_santa.users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES secret_santa.groups(id) ON DELETE CASCADE,
    UNIQUE (year, group_id, secret_santa_id, recipient_id)
);
