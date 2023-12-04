package service

import (
	"context"

	pb "accident-service/api/accident/v1"
	"accident-service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
func (s *AccidentService) ListAccident(context.Context, *emptypb.Empty) (*pb.ListAccidentReply, error) {
	accidents, total, err := s.uc.List(context.TODO())
	if err != nil {
		return nil, err
	}
	accidentsDTO := make([]*pb.AccidentReply, 0)
	for _, accident := range accidents {
		dto := &pb.AccidentReply{
			Id:        accident.Id,
			Name:      accident.Name,
			Lat:       float32(accident.Lat),
			Lon:       float32(accident.Lon),
			StartDate: timestamppb.New(accident.StartDate),
		}
		if accident.EndDate != nil {
			dto.EndDate = timestamppb.New(*accident.EndDate)
		}
		accidentsDTO = append(accidentsDTO, dto)
	}
	return &pb.ListAccidentReply{
		Total:     total,
		Accidents: accidentsDTO,
	}, nil
}
