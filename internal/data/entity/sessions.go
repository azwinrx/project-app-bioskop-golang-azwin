package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID  `db:"id"`
	UserID    int        `db:"user_id"`
	ExpiresAt time.Time  `db:"expires_at"`
	RevokedAt *time.Time `db:"revoked_at"`
	CreatedAt time.Time  `db:"created_at"`
}
