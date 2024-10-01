package Transactions

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	id   uuid.UUID
	time time.Time
	sum  int
}
