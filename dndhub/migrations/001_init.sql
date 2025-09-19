-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ================================================
-- Users (central DnDHub user accounts)
-- ================================================
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ================================================
-- Franchises (belongs to a user)
-- ================================================
CREATE TABLE franchises (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,

    funds_cents BIGINT NOT NULL DEFAULT 0,
    property_value_cents BIGINT NOT NULL DEFAULT 0,

    unskilled_workers INT NOT NULL DEFAULT 0,
    lowskilled_workers INT NOT NULL DEFAULT 0,
    highskilled_workers INT NOT NULL DEFAULT 0,

    cost_unskilled_cents INT NOT NULL DEFAULT 0,
    cost_lowskilled_cents INT NOT NULL DEFAULT 0,
    cost_highskilled_cents INT NOT NULL DEFAULT 0,

    revenue_modifier_bp INT NOT NULL DEFAULT 0,
    upkeep_modifier_bp INT NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ================================================
-- Unique Workers (RPG-style employees)
-- ================================================
CREATE TABLE unique_workers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    franchise_id UUID NOT NULL REFERENCES franchises(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    monthly_cost_cents INT NOT NULL DEFAULT 0,

    creativity INT NOT NULL DEFAULT 0,
    discipline INT NOT NULL DEFAULT 0,
    charisma INT NOT NULL DEFAULT 0,
    efficiency INT NOT NULL DEFAULT 0,
    exploration INT NOT NULL DEFAULT 0,

    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
