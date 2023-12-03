package service

import (
	"context"

	pb "accident-service/api/accident/v1"
	"accident-service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AccidentService struct {
	uc *biz.AccidentUseCase
	pb.UnimplementedAccidentServer
}

func NewAccidentService(uc *biz.AccidentUseCase) *AccidentService {
	return &AccidentService{
		uc: uc,
	}
}

func (s *AccidentService) CreateAccident(ctx context.Context, req *pb.CreateAccidentRequest) (*emptypb.Empty, error) {
	accident := &biz.Accident{
		Name:      req.Name,
		Lat:       float64(req.Lat),
		Lon:       float64(req.Lon),
		StartDate: req.StartDate.AsTime(),
	}
	if req.EndDate != nil {
		time := req.EndDate.AsTime()
		accident.EndDate = &time
	}

	err := s.uc.Create(ctx, accident)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
