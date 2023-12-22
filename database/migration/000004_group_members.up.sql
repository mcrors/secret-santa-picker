CREATE TABLE IF NOT EXISTS secret_santa.group_members (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL UNIQUE,
    user_id INT NOT NULL,
    group_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES secret_santa.users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES secret_santa.groups(id) ON DELETE CASCADE
);
