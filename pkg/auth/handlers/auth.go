package handlers

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/auth"
)

func DeleteUser(id string, auth pb.AuthServiceClient) (*pb.Response, error) {
	res, err := auth.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		Id: id,
	})

	if err != nil {
		return &pb.Response{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, err
	}

	return &pb.Response{
		Status: res.Status,
	}, nil
}

func UpdatePassword(user *models.PasswordUpdate, auth pb.AuthServiceClient) (*pb.Response, error) {
	res, err := auth.UpdatePassword(context.Background(), &pb.UpdatePasswordRequest{
		Id:          user.Id,
		Newpassword: user.NewPassword,
		Oldpassword: user.OldPassword,
	})

	if err != nil {
		return &pb.Response{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, err
	}

	return &pb.Response{
		Status: res.Status,
	}, nil
}
