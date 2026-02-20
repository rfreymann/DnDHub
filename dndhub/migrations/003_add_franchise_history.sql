CREATE TABLE franchise_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    franchise_id UUID NOT NULL REFERENCES franchises(id) ON DELETE CASCADE,
    revenue BIGINT NOT NULL DEFAULT 0,
    expenses BIGINT NOT NULL DEFAULT 0,
    profit BIGINT NOT NULL DEFAULT 0,
    roll INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
