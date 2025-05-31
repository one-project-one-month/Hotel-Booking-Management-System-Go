package events

import "github.com/google/uuid"

type FindByUserIdDto struct {
	UserID string
}

type FindByIdDto struct {
	ID uuid.UUID
}

type FindByAccountNumberDto struct {
	AccountNumber string
}
