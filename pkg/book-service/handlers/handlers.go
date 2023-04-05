package handlers

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/book"
)

func GetBookId(book *models.Book, cl pb.BookServiceClient) (*pb.GetBookIdResponse, error) {
	res, err := cl.GetBookId(context.Background(), &pb.GetBookIdRequest{
		Book: &pb.Book{
			Author: &pb.Author{
				Authorid: book.AuthorId,
				Name:     book.Name,
				Surname:  book.Surname,
			},
			Title:  book.Title,
			Bookid: book.BookId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetBookIdResponse{Id: res.Id}, nil
}
