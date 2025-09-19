package franchise

import (
	"time"

	"github.com/google/uuid"
)

type UniqueWorker struct {
	ID               uuid.UUID `db:"id" json:"id"`
	FranchiseID      uuid.UUID `db:"franchise_id" json:"franchise_id"`
	Name             string    `db:"name" json:"name"`
	MonthlyCostCents int       `db:"monthly_cost_cents" json:"monthly_cost_cents"`
	Creativity       int       `db:"creativity" json:"creativity"`
	Discipline       int       `db:"discipline" json:"discipline"`
	Charisma         int       `db:"charisma" json:"charisma"`
	Efficiency       int       `db:"efficiency" json:"efficiency"`
	Exploration      int       `db:"exploration" json:"exploration"`
	Notes            *string   `db:"notes" json:"notes"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
