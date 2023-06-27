package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/user"
)

func (h *Handler) GetAllReviews(ctx context.Context, req *pb.GetAllReviewsRequest) (*pb.GetAllReviewsResponse, error) {
	posts, err := h.se.GetReviewsByAllUsers(ctx)
	if err != nil {
		return &pb.GetAllReviewsResponse{
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

	return &pb.GetAllReviewsResponse{
		Feed: feed,
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	err := h.se.DeletePost(ctx, req.Postid, req.Userid)
	if err != nil {
		return &pb.DeletePostResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.DeletePostResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	newpost := models.Post{
		AuthorName:    req.Newpost.AuthorName,
		AuthorSurname: req.Newpost.AuthorSurname,
		Title:         req.Newpost.Title,
		Content:       req.Newpost.Content,
	}

	check := h.validate(&newpost)

	if !check {
		return &pb.CreatePostResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  "Field cant be empty",
			},
		}, errors.New("Field cant be empty")
	}

	err := h.se.CreatePost(ctx, &newpost, req.Id)
	if err != nil {
		return &pb.CreatePostResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.CreatePostResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := h.se.GetPost(ctx, req.Postid)
	if err != nil {
		return &pb.GetPostResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.GetPostResponse{
		Post: &pb.Post{
			AuthorName:    post.AuthorName,
			AuthorSurname: post.AuthorSurname,
			Title:         post.Title,
			Content:       post.Content,
		},
	}, nil
}

func (h *Handler) GetPostsForBook(ctx context.Context, req *pb.GetPostsForBookRequest) (*pb.GetPostsForBookResponse, error) {
	posts, err := h.se.GetPosts(ctx, req.Id)
	if err != nil {
		return &pb.GetPostsForBookResponse{
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

	return &pb.GetPostsForBookResponse{
		Feed: feed,
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (h *Handler) validate(post *models.Post) bool {
	if post.AuthorName == "" {
		return false
	}

	if post.AuthorSurname == "" {
		return false
	}

	if post.Content == "" {
		return false
	}

	if post.Title == "" {
		return false
	}

	return true
}
