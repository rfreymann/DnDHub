package franchise

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateFranchise(ctx context.Context, userID uuid.UUID, name string) (Franchise, error) {
	var f Franchise
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO franchises (user_id, name)
		VALUES ($1, $2)
		RETURNING id, user_id, name,
		          funds_cents, property_value_cents,
		          unskilled_workers, lowskilled_workers, highskilled_workers,
		          cost_unskilled_cents, cost_lowskilled_cents, cost_highskilled_cents,
		          revenue_modifier_bp, upkeep_modifier_bp,
		          created_at, updated_at
	`, userID, name).Scan(
		&f.ID, &f.UserID, &f.Name,
		&f.FundsCents, &f.PropertyValueCents,
		&f.UnskilledWorkers, &f.LowskilledWorkers, &f.HighskilledWorkers,
		&f.CostUnskilledCents, &f.CostLowskilledCents, &f.CostHighskilledCents,
		&f.RevenueModifierBP, &f.UpkeepModifierBP,
		&f.CreatedAt, &f.UpdatedAt,
	)
	return f, err
}

func (r *Repository) GetFranchisesByUser(ctx context.Context, userID uuid.UUID) ([]Franchise, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, name,
		       funds_cents, property_value_cents,
		       unskilled_workers, lowskilled_workers, highskilled_workers,
		       cost_unskilled_cents, cost_lowskilled_cents, cost_highskilled_cents,
		       revenue_modifier_bp, upkeep_modifier_bp,
		       created_at, updated_at
		FROM franchises
		WHERE user_id = $1
		ORDER BY created_at ASC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Franchise
	for rows.Next() {
		var f Franchise
		if err := rows.Scan(
			&f.ID, &f.UserID, &f.Name,
			&f.FundsCents, &f.PropertyValueCents,
			&f.UnskilledWorkers, &f.LowskilledWorkers, &f.HighskilledWorkers,
			&f.CostUnskilledCents, &f.CostLowskilledCents, &f.CostHighskilledCents,
			&f.RevenueModifierBP, &f.UpkeepModifierBP,
			&f.CreatedAt, &f.UpdatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, rows.Err()
}

func (r *Repository) GetFranchiseByID(ctx context.Context, userID, franchiseID uuid.UUID) (Franchise, error) {
	var f Franchise
	err := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, name,
		       funds_cents, property_value_cents,
		       unskilled_workers, lowskilled_workers, highskilled_workers,
		       cost_unskilled_cents, cost_lowskilled_cents, cost_highskilled_cents,
		       revenue_modifier_bp, upkeep_modifier_bp,
		       created_at, updated_at
		FROM franchises
		WHERE id = $1 AND user_id = $2
	`, franchiseID, userID).Scan(
		&f.ID, &f.UserID, &f.Name,
		&f.FundsCents, &f.PropertyValueCents,
		&f.UnskilledWorkers, &f.LowskilledWorkers, &f.HighskilledWorkers,
		&f.CostUnskilledCents, &f.CostLowskilledCents, &f.CostHighskilledCents,
		&f.RevenueModifierBP, &f.UpkeepModifierBP,
		&f.CreatedAt, &f.UpdatedAt,
	)
	return f, err
}

func (r *Repository) GetWorkersByFranchise(ctx context.Context, franchiseID uuid.UUID) ([]UniqueWorker, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, franchise_id, name, monthly_cost_cents,
		       creativity, discipline, charisma, efficiency, exploration,
		       notes, created_at, updated_at
		FROM unique_workers
		WHERE franchise_id = $1
		ORDER BY created_at ASC
	`, franchiseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []UniqueWorker
	for rows.Next() {
		var w UniqueWorker
		if err := rows.Scan(
			&w.ID, &w.FranchiseID, &w.Name, &w.MonthlyCostCents,
			&w.Creativity, &w.Discipline, &w.Charisma, &w.Efficiency, &w.Exploration,
			&w.Notes, &w.CreatedAt, &w.UpdatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, w)
	}
	return list, rows.Err()
}

func (r *Repository) UpdateFranchise(ctx context.Context, userID, franchiseID uuid.UUID, f Franchise) (Franchise, error) {
	var updated Franchise
	err := r.db.QueryRowContext(ctx, `
		UPDATE franchises
		SET name=$1,
		    funds_cents=$2,
		    property_value_cents=$3,
		    unskilled_workers=$4,
		    lowskilled_workers=$5,
		    highskilled_workers=$6,
		    cost_unskilled_cents=$7,
		    cost_lowskilled_cents=$8,
		    cost_highskilled_cents=$9,
		    revenue_modifier_bp=$10,
		    upkeep_modifier_bp=$11,
		    updated_at=now()
		WHERE id=$12 AND user_id=$13
		RETURNING id, user_id, name,
		          funds_cents, property_value_cents,
		          unskilled_workers, lowskilled_workers, highskilled_workers,
		          cost_unskilled_cents, cost_lowskilled_cents, cost_highskilled_cents,
		          revenue_modifier_bp, upkeep_modifier_bp,
		          created_at, updated_at
	`,
		f.Name, f.FundsCents, f.PropertyValueCents,
		f.UnskilledWorkers, f.LowskilledWorkers, f.HighskilledWorkers,
		f.CostUnskilledCents, f.CostLowskilledCents, f.CostHighskilledCents,
		f.RevenueModifierBP, f.UpkeepModifierBP,
		franchiseID, userID,
	).Scan(
		&updated.ID, &updated.UserID, &updated.Name,
		&updated.FundsCents, &updated.PropertyValueCents,
		&updated.UnskilledWorkers, &updated.LowskilledWorkers, &updated.HighskilledWorkers,
		&updated.CostUnskilledCents, &updated.CostLowskilledCents, &updated.CostHighskilledCents,
		&updated.RevenueModifierBP, &updated.UpkeepModifierBP,
		&updated.CreatedAt, &updated.UpdatedAt,
	)
	return updated, err
}

func (r *Repository) CreateWorker(ctx context.Context, franchiseID uuid.UUID, w UniqueWorker) (UniqueWorker, error) {
	var created UniqueWorker
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO unique_workers
		(franchise_id, name, monthly_cost_cents, creativity, discipline, charisma, efficiency, exploration, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, franchise_id, name, monthly_cost_cents,
		          creativity, discipline, charisma, efficiency, exploration,
		          notes, created_at, updated_at
	`,
		franchiseID, w.Name, w.MonthlyCostCents,
		w.Creativity, w.Discipline, w.Charisma, w.Efficiency, w.Exploration, w.Notes,
	).Scan(
		&created.ID, &created.FranchiseID, &created.Name, &created.MonthlyCostCents,
		&created.Creativity, &created.Discipline, &created.Charisma, &created.Efficiency, &created.Exploration,
		&created.Notes, &created.CreatedAt, &created.UpdatedAt,
	)
	return created, err
}

func (r *Repository) UpdateWorker(ctx context.Context, franchiseID, workerID uuid.UUID, w UniqueWorker) (UniqueWorker, error) {
	var updated UniqueWorker
	err := r.db.QueryRowContext(ctx, `
		UPDATE unique_workers
		SET name=$1,
		    monthly_cost_cents=$2,
		    creativity=$3,
		    discipline=$4,
		    charisma=$5,
		    efficiency=$6,
		    exploration=$7,
		    notes=$8,
		    updated_at=now()
		WHERE id=$9 AND franchise_id=$10
		RETURNING id, franchise_id, name, monthly_cost_cents,
		          creativity, discipline, charisma, efficiency, exploration,
		          notes, created_at, updated_at
	`,
		w.Name, w.MonthlyCostCents, w.Creativity, w.Discipline, w.Charisma,
		w.Efficiency, w.Exploration, w.Notes, workerID, franchiseID,
	).Scan(
		&updated.ID, &updated.FranchiseID, &updated.Name, &updated.MonthlyCostCents,
		&updated.Creativity, &updated.Discipline, &updated.Charisma, &updated.Efficiency, &updated.Exploration,
		&updated.Notes, &updated.CreatedAt, &updated.UpdatedAt,
	)
	return updated, err
}

func (r *Repository) DeleteWorker(ctx context.Context, franchiseID, workerID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM unique_workers
		WHERE id=$1 AND franchise_id=$2
	`, workerID, franchiseID)
	return err
}
