package Transactions

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	id   uuid.UUID
	time time.Time
	sum  int
	// FIXME: should we store user id inside Transaction model? Or create extend model UserTransaction where we would store trnsctn id and usr id together?
}
