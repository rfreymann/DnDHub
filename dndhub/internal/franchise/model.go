package franchise

import (
	"time"

	"github.com/google/uuid"
)

type Franchise struct {
	ID                   uuid.UUID `db:"id" json:"id"`
	UserID               uuid.UUID `db:"user_id" json:"user_id"`
	Name                 string    `db:"name" json:"name"`
	FundsCents           int64     `db:"funds_cents" json:"funds_cents"`
	PropertyValueCents   int64     `db:"property_value_cents" json:"property_value_cents"`
	UnskilledWorkers     int       `db:"unskilled_workers" json:"unskilled_workers"`
	LowskilledWorkers    int       `db:"lowskilled_workers" json:"lowskilled_workers"`
	HighskilledWorkers   int       `db:"highskilled_workers" json:"highskilled_workers"`
	CostUnskilledCents   int       `db:"cost_unskilled_cents" json:"cost_unskilled_cents"`
	CostLowskilledCents  int       `db:"cost_lowskilled_cents" json:"cost_lowskilled_cents"`
	CostHighskilledCents int       `db:"cost_highskilled_cents" json:"cost_highskilled_cents"`
	RevenueModifierBP    int       `db:"revenue_modifier_bp" json:"revenue_modifier_bp"`
	UpkeepModifierBP     int       `db:"upkeep_modifier_bp" json:"upkeep_modifier_bp"`
	CreatedAt            time.Time `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
}

type FranchiseWithWorkers struct {
	Franchise
	Workers []UniqueWorker `json:"workers"`
}
