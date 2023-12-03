package data

import (
	"accident-service/internal/biz"
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

func (m Accident) modelToResponse() *biz.Accident {
	return &biz.Accident{
		Id:        m.Id,
		Name:      m.Name,
		Lat:       m.Lat,
		Lon:       m.Lon,
		StartDate: m.StartDate,
		EndDate:   m.EndDate,
	}
}

type accidentRepo struct {
	data   *Data
	logger *log.Helper
}

func NewAccidentRepo(data *Data, logger log.Logger) biz.AccidentRepo {
	return &accidentRepo{data: data, logger: log.NewHelper(logger)}
}

// Create implements biz.AccidentRepo.
func (r *accidentRepo) Create(ctx context.Context, accident *biz.Accident) error {
	var accidentDB Accident
	accidentDB.Name = accident.Name
	accidentDB.Lat = accident.Lat
	accidentDB.Lon = accident.Lon
	accidentDB.StartDate = accident.StartDate
	accidentDB.EndDate = accident.EndDate
	if err := r.data.db.Create(&accidentDB).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements biz.AccidentRepo.
func (r *accidentRepo) Delete(ctx context.Context, id uint64) error {
	return r.data.db.Delete(&Accident{}, id).Error
}

// List implements biz.AccidentRepo.
func (r *accidentRepo) List(ctx context.Context) ([]*biz.Accident, int64, error) {
	var listDB []Accident
	localDB := r.data.db.Model(&Accident{})
	if err := localDB.Find(&listDB).Error; err != nil {
		return nil, 0, err
	}
	var count int64
	localDB.Count(&count)
	lists := make([]*biz.Accident, 0)
	for _, b := range listDB {
		lists = append(lists, b.modelToResponse())
	}
	return lists, count, nil
}

// Update implements biz.AccidentRepo.
func (r *accidentRepo) Update(ctx context.Context, accident *biz.Accident) error {
	var accidentDB Accident
	accidentDB.Name = accident.Name
	accidentDB.Lat = accident.Lat
	accidentDB.Lon = accident.Lon
	accidentDB.StartDate = accident.StartDate
	accidentDB.EndDate = accident.EndDate
	accidentDB.Id = accident.Id
	if err := r.data.db.Save(&accidentDB).Error; err != nil {
		return err
	}
	return nil
}
