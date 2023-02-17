package handlers

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/user"
)

func (h *Handler) CreateUser(ctx context.Context, uscl *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	if err := h.se.CreateUser(ctx, &models.UserModel{
		ID:       uscl.Id,
		Name:     uscl.Name,
		Username: uscl.Username,
		Email:    uscl.Email,
	}); err != nil {
		return &pb.CreateUserResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return &pb.CreateUserResponse{
		Status: http.StatusOK,
	}, nil
}

func (h *Handler) GetNewUsers(ctx context.Context, uscl *pb.GetNewUsersRequest) (*pb.GetNewUsersResponse, error) {
	lastUsers, err := h.se.GetNewUsers(ctx)

	var userList []*pb.User

	for _, users := range lastUsers {
		user := new(pb.User)
		user.Username = users.Username
		user.Id = users.Id
		userList = append(userList, user)
	}

	if err != nil {
		return &pb.GetNewUsersResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.GetNewUsersResponse{
		User: userList,
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, uscl *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, posts, subs, err := h.se.GetUser(ctx, uscl.Id)
	if err != nil {
		return &pb.GetUserResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	var subList []*pb.User

	for _, sub := range subs {
		subUser := new(pb.User)
		subUser.Name = sub.Name
		subUser.Username = sub.Username
		subUser.JoinDate = sub.JoinDate
		subUser.Status = sub.Status
		subUser.Email = sub.Email
		subList = append(subList, subUser)
	}

	var postList []*pb.Post

	for _, post := range posts {
		postProto := new(pb.Post)
		postProto.AuthorName = post.AuthorName
		postProto.AuthorSurname = post.AuthorSurname
		postProto.Content = post.Content
		postProto.Title = post.Title
		postList = append(postList, postProto)
	}

	return &pb.GetUserResponse{
		Post: postList,
		User: &pb.User{
			Username: user.Username,
			Name:     user.Name,
			Status:   user.Status,
			JoinDate: user.JoinDate,
			Email:    user.Email,
		},
		Subscriptions: subList,
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, uscl *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := models.UserUpdate{
		Name:     *uscl.Name,
		Username: *uscl.Username,
		Status:   *uscl.Status,
	}

	newpassword := models.PasswordUpdate{
		NewPassword: *uscl.Newpassword,
		OldPassword: *uscl.Oldpassword,
		Id:          uscl.Id,
	}
	
	if newpassword.NewPassword != "" {
		res, err := h.se.Client.UpdatePassword(&newpassword)
		if err != nil {
			return &pb.UpdateUserResponse{
				Response: &pb.Response{
					Status: res.Status,
					Error:  res.Error,
				},
			}, err
		}
	}

	err := h.se.UpdateUser(ctx, uscl.Id, &user)
	if err != nil {
		return &pb.UpdateUserResponse{
			Response: &pb.Response{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.UpdateUserResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, uscl *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.se.DeleteUser(ctx, uscl.Id)
	if err != nil {
		return &pb.DeleteUserResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	res, err := h.se.Client.DeleteUser(uscl.Id)
	if err != nil {
		return &pb.DeleteUserResponse{
			Response: &pb.Response{
				Status: res.Status,
				Error:  res.Error,
			},
		}, err
	}

	return &pb.DeleteUserResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}
