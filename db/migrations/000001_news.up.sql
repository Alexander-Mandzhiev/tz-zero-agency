CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password_hash BYTEA
);

CREATE TABLE IF NOT EXISTS news (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    title VARCHAR(255) UNIQUE,
    content TEXT
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    title VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS news_categories (
    news_id INT REFERENCES news (id) ON DELETE CASCADE,
    category_id INT REFERENCES categories (id) ON DELETE CASCADE,
    PRIMARY KEY (news_id, category_id)
);

CREATE INDEX idx_id ON users (id);

CREATE INDEX idx_email ON users (email);

CREATE INDEX idx_user_id_news ON news (user_id);

CREATE INDEX idx_title_news ON news (title);

CREATE INDEX idx_user_id_categories ON categories (user_id);

CREATE INDEX idx_title_categories ON categories (title);