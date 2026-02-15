CREATE TABLE IF NOT EXISTS public.members (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL UNIQUE,
    group_id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(75) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE
);
