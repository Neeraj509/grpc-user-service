package service

import (
	"context"
	"errors"
	"fmt"
	"user_service/models"
	"user_service/proto"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	users []models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: []models.User{
			{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{ID: 2, FName: "Alice", City: "New York", Phone: 9876543210, Height: 5.6, Married: false},
			{ID: 3, FName: "John", City: "Chicago", Phone: 5551234567, Height: 6.0, Married: true},
			{ID: 4, FName: "Emily", City: "San Francisco", Phone: 1112223333, Height: 5.4, Married: false},
			// Add more sample users as needed
		},
	}
}

func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	fmt.Println("request", req)
	for _, user := range s.users {
		if user.ID == req.Id {
			return &proto.GetUserResponse{User: &proto.User{
				Id:      user.ID,
				Fname:   user.FName,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			}}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserService) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	var users []*proto.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.ID == id {
				users = append(users, &proto.User{
					Id:      user.ID,
					Fname:   user.FName,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
			}
		}
	}
	return &proto.GetUsersResponse{Users: users}, nil
}

func (s *UserService) SearchUsers(ctx context.Context, req *proto.SearchUsersRequest) (*proto.SearchUsersResponse, error) {
	var users []*proto.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(!req.Married || user.Married == req.Married) {
			users = append(users, &proto.User{
				Id:      user.ID,
				Fname:   user.FName,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}
	return &proto.SearchUsersResponse{Users: users}, nil
}
