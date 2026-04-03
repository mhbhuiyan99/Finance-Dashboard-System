-- +migrate Up
CREATE TABLE financial_records (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount      NUMERIC(15, 2) NOT NULL CHECK (amount > 0),
    type        VARCHAR(10) NOT NULL CHECK (type IN ('income', 'expense')),
    category    VARCHAR(100) NOT NULL,
    date        DATE NOT NULL,
    notes       TEXT DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE INDEX idx_records_user_id ON financial_records(user_id);
CREATE INDEX idx_records_type ON financial_records(type);
CREATE INDEX idx_records_category ON financial_records(category);
CREATE INDEX idx_records_date ON financial_records(date);
CREATE INDEX idx_records_deleted_at ON financial_records(deleted_at);

