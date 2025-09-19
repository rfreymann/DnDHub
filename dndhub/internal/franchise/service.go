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
