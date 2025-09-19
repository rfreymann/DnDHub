-- ================================================
-- Sessions (for login tokens / JWT tracking)
-- ================================================

CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    -- opaque session token or JWT ID (jti)
    token TEXT UNIQUE NOT NULL,

    -- expiry date for the session
    expires_at TIMESTAMPTZ NOT NULL,

    -- optional: user agent or IP for session tracking
    client_info TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- index for quick lookup of active sessions
CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_token_expires ON sessions(token, expires_at);
