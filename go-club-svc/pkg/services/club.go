package services

import (
	"club/pkg/db"
	"club/pkg/models"
	"club/pkg/pb"
	"context"
	"net/http"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateClub(ctx context.Context, req *pb.CreateClubRequest) (*pb.CreateClubResponse, error) {
	var club models.Club

	club.Name = req.Name
	club.Year = req.Year
	club.Titles = req.Trophy

	if result := s.H.DB.Create(&club); result.Error != nil {
		return &pb.CreateClubResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.CreateClubResponse{
		Status: http.StatusCreated,
		Id:    club.Id,
	}, nil
}

func (s *Server) FindClub(ctx context.Context, req *pb.FindClubRequest) (*pb.FindClubResponse, error) {
	var club models.Club

	if result := s.H.DB.First(&club, req.Id); result.Error != nil {
		return &pb.FindClubResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindClubData{
		Id: club.Id,
		Name: club.Name,
		Year: club.Year,
		Trophy: club.Titles,
	}

	return &pb.FindClubResponse{
		Status: http.StatusOK,
		Data:   data,
	},nil
}