package handlers

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/user"
)

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	if err := h.se.CreateUser(ctx, &models.UserModel{
		ID:       req.Id,
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
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

func (h *Handler) GetNewUsers(ctx context.Context, req *pb.GetNewUsersRequest) (*pb.GetNewUsersResponse, error) {
	lastUsers, err := h.se.GetNewUsers(ctx, req.Id)

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

func (h *Handler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, posts, subs, err := h.se.GetUser(ctx, req.Id)
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

func (h *Handler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := models.UserUpdate{
		Name:     *req.Name,
		Username: *req.Username,
		Status:   *req.Status,
	}

	newpassword := models.PasswordUpdate{
		NewPassword: *req.Newpassword,
		OldPassword: *req.Oldpassword,
		Id:          req.Id,
	}

	if newpassword.NewPassword != "" {
		res, err := h.se.AuthClient.UpdatePassword(&newpassword)
		if err != nil {
			return &pb.UpdateUserResponse{
				Response: &pb.Response{
					Status: res.Status,
					Error:  res.Error,
				},
			}, err
		}

		return &pb.UpdateUserResponse{
			Response: &pb.Response{
				Status: http.StatusOK,
			},
		}, nil
	}

	err := h.se.UpdateUser(ctx, req.Id, &user)
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

func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.se.DeleteUser(ctx, req.Id)
	if err != nil {
		return &pb.DeleteUserResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	res, err := h.se.AuthClient.DeleteUser(req.Id)
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

func (h *Handler) AddNewSubscription(ctx context.Context, req *pb.AddNewSubscriptionRequest) (*pb.AddNewSubscriptionResponse, error) {
	err := h.se.AddNewSubscription(ctx, req.Userid, req.Id)
	if err != nil {
		return &pb.AddNewSubscriptionResponse{
			Response: &pb.Response{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.AddNewSubscriptionResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) DeleteSubscription(ctx context.Context, req *pb.DeleteSubscriptionRequest) (*pb.DeleteSubscriptionResponse, error) {
	err := h.se.DeleteOneSubscription(ctx, req.Userid, req.Id)
	if err != nil {
		return &pb.DeleteSubscriptionResponse{
			Response: &pb.Response{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.DeleteSubscriptionResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) GetMyFeed(ctx context.Context, req *pb.GetMyFeedRequest) (*pb.GetMyFeedResponse, error) {
	posts, err := h.se.GetMyFeed(ctx, req.Id)
	if err != nil {
		return &pb.GetMyFeedResponse{
			Response: &pb.Response{
				Status: http.StatusOK,
				Error:  err.Error(),
			},
		}, err
	}

	var feed []*pb.Feed

	for _, post := range posts {
		feed_one := pb.Feed{}
		feed_one.Username = post.Username
		feed_one.Status = post.Status
		feed_one.Date = post.CreateDate
		feed_one.AuthorName = post.AuthorName
		feed_one.AuthorSurname = post.AuthorSurname
		feed_one.Title = post.Title
		feed_one.Content = post.Content
		feed_one.Userid = post.UserId
		feed = append(feed, &feed_one)
	}

	return &pb.GetMyFeedResponse{
		Feed: feed,
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) CheckMySubs(ctx context.Context, req *pb.CheckMySubsRequest) (*pb.CheckMySubsResponse, error) {
	check, err := h.se.CheckSubs(ctx, req.Myid, req.Userid)
	if err != nil {
		return &pb.CheckMySubsResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}
	return &pb.CheckMySubsResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		Boolcheck: check,
	}, nil
}

func (h *Handler) GetMySubscriptions(ctx context.Context, req *pb.GetMySubscriptionsRequest) (*pb.GetMySubscriptionsResponse, error) {
	subs, err := h.se.GetMySubs(ctx, req.Id)
	if err != nil {
		return &pb.GetMySubscriptionsResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	var users []*pb.User

	for _, sub := range subs {
		user_one := pb.User{}
		user_one.Name = sub.Name
		user_one.Username = sub.Username
		user_one.Status = sub.Status
		user_one.Email = sub.Email
		user_one.Id = sub.ID
		user_one.JoinDate = sub.JoinDate
		users = append(users, &user_one)
	}

	return &pb.GetMySubscriptionsResponse{
		User: users,
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}
