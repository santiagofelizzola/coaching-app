package models

import (
    "time"
)

type Session struct {
    BaseModel
    UserID    int
    Token     string
    ExpiresAt time.Time

    User User
}
