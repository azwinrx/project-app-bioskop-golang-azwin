-- Migration: Add sessions table for session-based authentication
-- Run this SQL to create the sessions table in your database

CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id INTEGER NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at);
CREATE INDEX IF NOT EXISTS idx_sessions_revoked_at ON sessions(revoked_at);

-- Useful queries for session management:
-- 
-- 1. Check active sessions:
-- SELECT * FROM sessions WHERE revoked_at IS NULL AND expires_at > NOW();
--
-- 2. Revoke all user sessions:
-- UPDATE sessions SET revoked_at = NOW() WHERE user_id = ? AND revoked_at IS NULL;
--
-- 3. Clean up expired sessions (maintenance):
-- DELETE FROM sessions WHERE expires_at < NOW() - INTERVAL '30 days';
