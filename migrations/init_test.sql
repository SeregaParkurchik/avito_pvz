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