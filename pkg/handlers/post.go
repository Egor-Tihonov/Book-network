package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/user"
)

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

func (h *Handler) GetLastPosts(ctx context.Context, req *pb.GetLastPostsRequest) (*pb.GetLastPostsResponse, error) {
	posts, err := h.se.GetLastPosts(ctx, req.Id)

	var postsList []*pb.LastPost

	for _, post := range posts {
		newpost := pb.LastPost{}
		newpost.Id = post.PostId
		newpost.Title = post.Title
		postsList = append(postsList, &newpost)
	}

	if err != nil {
		return &pb.GetLastPostsResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
		}, err
	}

	return &pb.GetLastPostsResponse{
		Posts: postsList,
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
