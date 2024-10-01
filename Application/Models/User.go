package Models

import "time"

type User struct {
	userId                    int
	login                     string
	password                  string
	budget                    int // income - mandatory expenses
	currentBalance            int // income - mandatory expenses - other expenses
	calculationExpirationDate time.Time
}
