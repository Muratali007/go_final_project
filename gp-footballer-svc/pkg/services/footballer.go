package services

import (
	"context"
	"football/pkg/db"
	"football/pkg/models"
	"football/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateFootballer(ctx context.Context, req *pb.CreateFootballerRequest) (*pb.CreateFootballerResponse,error) {
	var footballer models.Footballer

	footballer.Name =req.Name
	footballer.Price = req.Price
	footballer.Club = req.Club

	if result := s.H.DB.Create(&footballer); result.Error != nil {
		return &pb.CreateFootballerResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateFootballerResponse{
		Status: http.StatusCreated,
		Id:     footballer.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest)(*pb.FindOneResponse,error) {
	var footballer models.Footballer

	if result := s.H.DB.First(&footballer, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data :=&pb.FindOneData{
		Id: footballer.Id,
		Price: footballer.Price,
		Name: footballer.Name,
		Club: footballer.Club,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DeleteFootballer(ctx context.Context, req *pb.DeleteFootballerRequest)(*pb.DeleteFootballerResponse,error) {
	var footballer models.Footballer

	if result := s.H.DB.Delete(&footballer, req.Id); result.Error != nil {
		return &pb.DeleteFootballerResponse{
			Status: http.StatusInternalServerError,
			Error: result.Error.Error(),
		},nil
	}

	return &pb.DeleteFootballerResponse{
		Status: http.StatusOK,
		Error: "Deleted footballer",
	},nil
}