package service

import (
	"context"
	"testing"
	"user_service/proto"
)

func TestGetUser(t *testing.T) {
	svc := NewUserService()

	req := &proto.GetUserRequest{Id: 1}
	res, err := svc.GetUser(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.User.Id != 1 {
		t.Fatalf("expected user id 1, got %v", res.User.Id)
	}
}

func TestSearchUsers(t *testing.T) {
	svc := NewUserService()

	req := &proto.SearchUsersRequest{City: "LA"}
	res, err := svc.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(res.Users) == 0 {
		t.Fatalf("expected at least one user, got %v", len(res.Users))
	}
}
