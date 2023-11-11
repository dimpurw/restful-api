package service

import (
	"context"
	"restful-api/data/request"
	"restful-api/data/response"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	findById(ctx context.Context, bookId int) response.BookResponse
	findAll(ctx context.Context) []response.BookResponse
}