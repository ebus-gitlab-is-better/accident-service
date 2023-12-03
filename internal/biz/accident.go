package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Accident struct {
	Id   uint64 `gorm:"primaryKey"`
	Name string
	// Shape     string
	Lat       float64
	Lon       float64
	StartDate time.Time
	EndDate   *time.Time
}

type AccidentRepo interface {
	Create(context.Context, *Accident) error
	Update(context.Context, *Accident) error
	List(context.Context) ([]*Accident, int64, error)
	Delete(context.Context, uint64) error
}

type AccidentUseCase struct {
	repo   AccidentRepo
	logger *log.Helper
}

func NewAccidentRepo(repo AccidentRepo, logger log.Logger) *AccidentUseCase {
	return &AccidentUseCase{repo: repo, logger: log.NewHelper(logger)}
}

func (uc *AccidentUseCase) Create(ctx context.Context, accident *Accident) error {
	return uc.repo.Create(ctx, accident)
}

func (uc *AccidentUseCase) Update(ctx context.Context, accident *Accident) error {
	return uc.repo.Update(ctx, accident)
}

func (uc *AccidentUseCase) List(ctx context.Context) ([]*Accident, int64, error) {
	return uc.repo.List(ctx)
}

func (uc *AccidentUseCase) Delete(ctx context.Context, id uint64) error {
	return uc.repo.Delete(ctx, id)
}
