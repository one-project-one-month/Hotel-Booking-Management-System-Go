package events

import "github.com/google/uuid"

type FindByUserIdDto struct {
	UserID string
}

type FindByIdDto struct {
	ID uuid.UUID
}

type FindByEmailDto struct {
	Email string
}

type FindByAccountNumberDto struct {
	AccountNumber string
}
