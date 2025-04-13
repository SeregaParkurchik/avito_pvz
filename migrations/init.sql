-- Active: 1744531808310@@localhost@54321@avitopvz
-- Таблица Users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL -- moderator, employee
);

-- Таблица PVZ (Пункты выдачи заказов)
CREATE TABLE pvz (
    id UUID PRIMARY KEY,
    registration_date TIMESTAMPTZ NOT NULL,
    city VARCHAR(255) NOT NULL
);

-- Таблица Acceptances (Приемки товаров)
CREATE TABLE acceptances (
    id UUID PRIMARY KEY,
    date_time TIMESTAMPTZ NOT NULL,
    pvz_id UUID REFERENCES pvz(id),
    status VARCHAR(50) NOT NULL -- in_progress, close
);

-- Таблица Products (Товары)
CREATE TABLE products (
    id UUID PRIMARY KEY,
    date_time TIMESTAMPTZ NOT NULL,
    type VARCHAR(255) NOT NULL,
    acceptance_id UUID REFERENCES acceptances(id),
    pvz_id UUID REFERENCES pvz(id)
);

-- Индексы для таблицы pvz
CREATE INDEX idx_pvz_id ON pvz(id); 
CREATE INDEX idx_pvz_registration_date ON pvz(registration_date);
-- Индексы для таблицы products
CREATE INDEX idx_products_date_time ON products(date_time);
CREATE INDEX idx_products_acceptance_id ON products(acceptance_id);
CREATE INDEX idx_products_pvz_id ON products(pvz_id);
-- Индексы для таблицы acceptances
CREATE INDEX idx_acceptances_date_time ON acceptances(date_time);
CREATE INDEX idx_acceptances_pvz_id ON acceptances(pvz_id);
CREATE INDEX idx_acceptances_status ON acceptances(status);
-- Индексы для таблицы users
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_role ON users (role);