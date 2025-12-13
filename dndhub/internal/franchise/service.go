package franchise

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateFranchise(ctx context.Context, userID uuid.UUID, name string) (Franchise, error) {
	return s.repo.CreateFranchise(ctx, userID, name)
}

func (s *Service) ListFranchises(ctx context.Context, userID uuid.UUID) ([]Franchise, error) {
	return s.repo.GetFranchisesByUser(ctx, userID)
}

func (s *Service) GetFranchise(ctx context.Context, userID, franchiseID uuid.UUID) (Franchise, error) {
	return s.repo.GetFranchiseByID(ctx, userID, franchiseID)
}

func (s *Service) GetFranchiseWithWorkers(ctx context.Context, userID, franchiseID uuid.UUID) (FranchiseWithWorkers, error) {
	f, err := s.repo.GetFranchiseByID(ctx, userID, franchiseID)
	if err != nil {
		return FranchiseWithWorkers{}, err
	}

	workers, err := s.repo.GetWorkersByFranchise(ctx, franchiseID)
	if err != nil {
		return FranchiseWithWorkers{}, err
	}

	fw := FranchiseWithWorkers{
		Franchise: f,
		Workers:   workers,
	}

	if fw.Workers == nil {
		fw.Workers = []UniqueWorker{}
	}

	return fw, nil
}

func (s *Service) UpdateFranchise(ctx context.Context, userID, franchiseID uuid.UUID, f Franchise) (Franchise, error) {
	return s.repo.UpdateFranchise(ctx, userID, franchiseID, f)
}

func (s *Service) CreateWorker(ctx context.Context, franchiseID uuid.UUID, w UniqueWorker) (UniqueWorker, error) {
	return s.repo.CreateWorker(ctx, franchiseID, w)
}

func (s *Service) UpdateWorker(ctx context.Context, franchiseID, workerID uuid.UUID, w UniqueWorker) (UniqueWorker, error) {
	return s.repo.UpdateWorker(ctx, franchiseID, workerID, w)
}

func (s *Service) DeleteWorker(ctx context.Context, franchiseID, workerID uuid.UUID) error {
	return s.repo.DeleteWorker(ctx, franchiseID, workerID)
}

type ActivityType string

const (
	ActivityMarketing     ActivityType = "marketing"
	ActivityAccounting    ActivityType = "accounting"
	ActivityRestructuring ActivityType = "restructuring"
)

type ActivityRequest struct {
	Type     ActivityType `json:"type"`
	WorkerID uuid.UUID    `json:"worker_id"`
}

type SimulationResult struct {
	PreviousFunds     int64   `json:"previous_funds"`
	NewFunds          int64   `json:"new_funds"`
	Profit            int64   `json:"profit"`
	Revenue           int64   `json:"revenue"`
	Expenses          int64   `json:"expenses"`
	Message           string  `json:"message"`
	AppliedActivities []string
	Roll              int     `json:"roll"`
	Modifier          int     `json:"modifier"`
	Description       string  `json:"description"`
}

func d100() int {
	return 1 + int(uuid.New().ID()%100) // Simple deterministic-ish rand for now, or use math/rand
}

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s *Service) SimulateFranchise(ctx context.Context, userID, franchiseID uuid.UUID, activities []ActivityRequest) (SimulationResult, error) {
	// 1. Fetch Franchise and Workers
	fw, err := s.GetFranchiseWithWorkers(ctx, userID, franchiseID)
	if err != nil {
		return SimulationResult{}, err
	}

	// 2. Validate Activities
	if len(activities) > 2 {
		return SimulationResult{}, fmt.Errorf("maximum 2 activities allowed")
	}

	workerMap := make(map[uuid.UUID]UniqueWorker)
	for _, w := range fw.Workers {
		workerMap[w.ID] = w
	}

	usedWorkers := make(map[uuid.UUID]bool)
	var validActivities []ActivityRequest

	for _, act := range activities {
		if _, exists := workerMap[act.WorkerID]; !exists {
			return SimulationResult{}, fmt.Errorf("worker %s does not belong to franchise", act.WorkerID)
		}
		if usedWorkers[act.WorkerID] {
			return SimulationResult{}, fmt.Errorf("worker %s cannot perform multiple activities", act.WorkerID)
		}
		usedWorkers[act.WorkerID] = true
		validActivities = append(validActivities, act)
	}


	// 3. Base Financials Calculation
	
	// Base Revenue: Property Value * 15%
	landRevenue := int64(float64(fw.PropertyValueCents) * 0.15)
	landUpkeep := int64(float64(fw.PropertyValueCents) * 0.05)
	// Base Upkeep: Sum of workers
	workerExpenses := int64(fw.UnskilledWorkers) * int64(fw.CostUnskilledCents) +
		int64(fw.LowskilledWorkers) * int64(fw.CostLowskilledCents) +
		int64(fw.HighskilledWorkers) * int64(fw.CostHighskilledCents)
	for _, w := range fw.Workers {
		workerExpenses += int64(w.MonthlyCostCents)
	}

	baseRevenue := int64(landRevenue + workerExpenses)
	baseUpkeep := int64(landUpkeep + workerExpenses)

	// We update the modifiers FIRST, then use them in the calculation.
	// Usually implies the activity *this month* counts for *this month*.
	
	var appliedMessages []string
	
	// Temporary Accounting Reduction
	var accountingReductionPct float64 = 0.0

	for _, act := range validActivities {
		worker := workerMap[act.WorkerID]
		switch act.Type {
		case ActivityMarketing:
			addedVal := worker.Charisma * baseRevenue * 0.01
			fw.Franchise.RevenueModifierBP += addedVal
			appliedMessages = append(appliedMessages, fmt.Sprintf("Marketing by %s: +%d Revenue Mod", worker.Name, addedVal))
		
		case ActivityRestructuring:
			addedVal := worker.Discipline * baseUpkeep * 0.01
			fw.Franchise.UpkeepModifierBP += addedVal
			appliedMessages = append(appliedMessages, fmt.Sprintf("Restructuring by %s: +%d Cost Reduction Mod", worker.Name, addedVal))
			
		case ActivityAccounting:
			// Efficiency * 2% reduction (Temporary for this month)
			accountingReductionPct += float64(worker.Efficiency) * 0.02
			appliedMessages = append(appliedMessages, fmt.Sprintf("Accounting by %s: -%.0f%% Expenses (Temporary)", worker.Name, float64(worker.Efficiency)*2))
		}
	}



	// 4. Modifier Application
	// Apply Modifiers
	// Revenue = Base + Modifier
	totalRevenueBasis := baseRevenue + int64(fw.RevenueModifierBP)
	
	// Upkeep = Base - Modifier (Deducted from costs)
	if accountingReductionPct > 1.0 { accountingReductionPct = 1.0 }
	baseUpkeepAfterAccounting := float64(baseUpkeep) * (1.0 - accountingReductionPct)
	
	totalUpkeepBasis := int64(baseUpkeepAfterAccounting) - int64(fw.UpkeepModifierBP)
	if totalUpkeepBasis < 0 { totalUpkeepBasis = 0 } // Costs shouldn't be negative (unless that's a grant?)

	// 5. Roll Logic (d100)
	roll := rand.Intn(100) + 1
	modifiedRoll := roll 
	
	type Range struct {
		Low, High int
		Desc      string
		Mult      float64
	}
	ranges := []Range{
		{1, 10,  "Katastrophaler Monat", 0.00},
		{11, 20, "Großer Verlust",       0.50},
		{21, 30, "Kleiner Verlust",      0.75},
		{31, 40, "Unter den Erwartungen",1.00},
		{41, 60, "Durchschnittlich",     1.25},
		{61, 70, "Guter Monat",          1.50},
		{71, 80, "Sehr guter Monat",     1.75},
		{81, 90, "Fantastischer Monat",  2.00},
		{91, 100,"Legendärer Erfolg",    2.25},
	}

	var selectedRange Range
	found := false
	for _, r := range ranges {
		if modifiedRoll >= r.Low && modifiedRoll <= r.High {
			selectedRange = r
			found = true
			break
		}
	}
	if !found {
		// Should not happen
		return SimulationResult{}, fmt.Errorf("roll %d did not match any tier", modifiedRoll)
	}

	// Profit = (Revenue * Multiplier) - Expenses
	// Note: previous logic applied multiplier to RevenueBasis.
	effectiveRevenue := int64(float64(totalRevenueBasis) * selectedRange.Mult)
	profit := effectiveRevenue - totalUpkeepBasis
	
	// Update Funds
	originalFunds := fw.FundsCents
	fw.Franchise.FundsCents += profit

	// 6. Decay Modifiers (x0.9)
	fw.Franchise.RevenueModifierBP = int(float64(fw.Franchise.RevenueModifierBP) * 0.9)
	fw.Franchise.UpkeepModifierBP = int(float64(fw.Franchise.UpkeepModifierBP) * 0.9)

	// Persist
	_, err = s.repo.UpdateFranchise(ctx, userID, franchiseID, fw.Franchise)
	if err != nil {
		return SimulationResult{}, err
	}

	return SimulationResult{
		PreviousFunds:     originalFunds,
		NewFunds:          fw.Franchise.FundsCents,
		Profit:            profit,
		Revenue:           effectiveRevenue,
		Expenses:          totalUpkeepBasis,
		Message:           fmt.Sprintf("Simulation: %s (Roll: %d)", selectedRange.Desc, modifiedRoll),
		AppliedActivities: appliedMessages,
		Roll:              roll,
		Modifier:          0, // No roll modifier currently
		Description:       selectedRange.Desc,
	}, nil
}
