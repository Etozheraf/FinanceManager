package Users

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id                        uuid.UUID
	login                     string
	password                  string
	budget                    int // income - mandatory expenses
	currentBalance            int // income - mandatory expenses - other expenses
	calculationExpirationDate time.Time
}
