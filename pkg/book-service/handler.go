package bookservice

import (
	"github.com/Egor-Tihonov/Book-network/pkg/book-service/handlers"
	"github.com/Egor-Tihonov/Book-network/pkg/config"
	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/book"
)

func RegisterHandlers(conf config.Config) *ServiceClient {
	svc := ServiceClient{
		Client: InitBookClient(&conf),
	}

	return &svc
}

func (s *ServiceClient) GetBook(book *models.Book) (*pb.GetBookIdResponse, error) {
	return handlers.GetBookId(book, s.Client)
}
